package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"oauth/auth"
	"oauth/config"
	_ "oauth/database"
	"oauth/database/bean"
	"oauth/database/iredis"
	"oauth/database/schema"
	"oauth/logger"
	"strings"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
)

var (
	cookieNameForSessionID = "mgtv-oauth-sessionid"
	session                = sessions.New(sessions.Config{
		Cookie: cookieNameForSessionID,
		//Expires: 45 * time.Minute, // <=0 means unlimited life
	})
)

type accountForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type appForm struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Callback string `json:"callback"`
}

type authForm struct {
	Data string `json:"data"`
}

type Scope struct {
	Type    string
	Name    string
	Actions []string
}

type AuthScope struct {
	Timestamp int64 `json:"timestamp"`
	Scope     Scope `json:"scope"`
}

type ResourceAccountForm struct {
	Timestamp int64  `json:"timestamp"`
	Token     string `json:"token"`
}

func GetApp() *iris.Application {
	app := iris.New()

	//conf := config.Get()
	app.Post("/authorize", func(ctx iris.Context) {
		clientID := ctx.GetHeader("client_id")
		account := ctx.GetHeader("account")
		account = strings.TrimSpace(account)
		clientID = strings.TrimSpace(clientID)
		if clientID == "" {
			ctx.StatusCode(406)
			return
		}
		if account == "" {
			ctx.StatusCode(401)
			return
		}
		//TODO 校验 account 是否存在数据库，是否处于正常状态

		body, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			logger.Debug(err)
			ctx.StatusCode(500)
			return
		}
		defer ctx.Request().Body.Close()
		decryptBody, err := auth.DecryptBody(clientID, body)
		if err != nil {
			logger.Debug(err)
			ctx.StatusCode(500)
			return
		}
		scope := Scope{}
		if err := json.Unmarshal([]byte(decryptBody), &scope); err != nil {
			logger.Debug(err)
			ctx.StatusCode(500)
			return
		}

		log.Println(fmt.Sprintf("权限请请求 %s : %s : %s : %s", clientID, account, scope.Name, scope.Type))

		authScope := AuthScope{
			Timestamp: time.Now().UnixNano(),
			Scope:     scope,
		}

		encryptAuthScope, err := auth.EncryptBody(clientID, authScope)
		if err != nil {
			logger.Debug(err)
			ctx.StatusCode(500)
			return
		}
		ctx.StatusCode(200)
		ctx.WriteString(encryptAuthScope)
	})

	app.Get("/authorize", func(ctx iris.Context) {
		clientID := ctx.URLParam("client_id")
		clientID = strings.TrimSpace(clientID)
		if clientID == "" {
			ctx.StatusCode(406)
			return
		}
		now := time.Now().UnixNano()
		if timestamp, err := ctx.URLParamInt64("timestamp"); err != nil {
			ctx.StatusCode(406)
			return
			//与服务器误差小于5分钟 东八区
		} else if math.Abs(float64(now-timestamp)) > 1000*5*60 {
			ctx.StatusCode(406)
			return
		}

		//是否存在私有key
		appPKKey := fmt.Sprintf("app:pk:%s", clientID)

		if !iredis.Exist(appPKKey) {
			ctx.StatusCode(406)
			return
		}

		sess := session.Start(ctx)
		//用户是否已登陆
		if userAuthorized, err := sess.GetBoolean("user-authorized"); err != nil || !userAuthorized {
			ctx.ServeFile("static/login.html", false)
			return
		}

		privateKey, err := iredis.Get(appPKKey)
		if err != nil {
			ctx.StatusCode(500)
			ctx.WriteString(err.Error())
			return
		}
		username := sess.GetString("username")
		if username == "" {
			ctx.ServeFile("static/login.html", false)
			return
		}

		token, err := auth.CreateResourceToken(clientID, username, privateKey)
		if err != nil {
			ctx.StatusCode(500)
			ctx.WriteString(err.Error())
			return
		}
		cbURL, err := iredis.Get(fmt.Sprintf("app:cb:%s", clientID))
		if err != nil {
			ctx.StatusCode(500)
			ctx.WriteString(err.Error())
			return
		}
		callback := fmt.Sprintf("%s?token=%s", cbURL, token)
		ctx.StatusCode(301)
		ctx.Header("Location", callback)
	})

	app.Post("/user-registry", func(ctx context.Context) {
		account := accountForm{}
		ctx.ReadJSON(&account)
		username := strings.TrimSpace(account.Username)
		password := strings.TrimSpace(account.Password)
		if username == "" || password == "" {
			ctx.StatusCode(iris.StatusNotAcceptable)
			ctx.WriteString("用户名或密码不能为空")
			return
		}
		if err := bean.RegisterUser(username, password); err != nil {
			ctx.StatusCode(iris.StatusNotAcceptable)
			ctx.WriteString(err.Error())
		} else {
			ctx.StatusCode(200)
			ctx.WriteString("注册成功！")
		}
	})

	app.Post("/app-registry", func(ctx context.Context) {
		form := appForm{}
		ctx.ReadJSON(&form)
		username := strings.TrimSpace(form.Name)
		password := strings.TrimSpace(form.Password)
		if username == "" || password == "" {
			ctx.StatusCode(iris.StatusNotAcceptable)
			ctx.WriteString("用户名或密码不能为空")
			return
		}
		app := schema.Application{
			Name:     username,
			Password: password,
			Callback: form.Callback,
		}
		if err := bean.RegisterAppliction(&app); err != nil {
			ctx.StatusCode(iris.StatusNotAcceptable)
			ctx.WriteString(err.Error())
		} else {
			ctx.StatusCode(200)
			app.Password = "*******"
			iredis.Set(fmt.Sprintf("app:pk:%s", app.ClientID), app.PrivateKey)
			iredis.Set(fmt.Sprintf("app:cb:%s", app.ClientID), app.Callback)
			ctx.JSON(app)
		}
	})

	ResourceAPI := app.Party("/resource")

	ResourceAPI.Post("/account", func(ctx iris.Context) {
		clientID := ctx.GetHeader("client_id")
		clientID = strings.TrimSpace(clientID)
		if clientID == "" {
			ctx.StatusCode(406)
			return
		}

		body, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			logger.Debug(err)
			ctx.StatusCode(500)
			return
		}
		defer ctx.Request().Body.Close()

		postStr, err := auth.DecryptBody(clientID, body)

		if err != nil {
			ctx.StatusCode(406)
			return
		}

		postData := ResourceAccountForm{}

		if err := json.Unmarshal([]byte(postStr), &postData); err != nil {
			ctx.StatusCode(500)
			return
		}

		if math.Abs(float64(time.Now().UnixNano()-postData.Timestamp)) > 1000*5*60 {
			logger.Debug("时间戳超时")
			ctx.StatusCode(406)
			return
		}
		token := postData.Token
		username, err := auth.GetResourceToken(clientID, token)
		if err != nil {
			logger.Debug(err)
			ctx.StatusCode(500)
			return
		}
		result := map[string]interface{}{
			"timestamp": time.Now().UnixNano(),
			"username":  username,
		}
		encryptBody, err := auth.EncryptBody(clientID, result)
		if err != nil {
			logger.Debug(err)
			ctx.StatusCode(500)
			return
		}
		ctx.WriteString(encryptBody)
	})

	ResourceAPI.Post("/login", func(ctx iris.Context) {
		account := accountForm{}
		ctx.ReadJSON(&account)
		username := strings.TrimSpace(account.Username)
		password := strings.TrimSpace(account.Password)
		if username == "" || password == "" {
			ctx.StatusCode(403)
			return
		}

		if _, exist, err := bean.FindUser(username, password); err != nil {
			ctx.StatusCode(500)
			ctx.WriteString(err.Error())
		} else if !exist {
			ctx.StatusCode(403)
		} else {
			sess := session.Start(ctx)
			sess.Set("user-authorized", true)
			sess.Set("username", username)
			ctx.StatusCode(200)
		}
	})

	return app
}

func main() {
	app := GetApp()
	app.Run(iris.Addr(":"+config.Get().Port), iris.WithoutVersionChecker)
}
