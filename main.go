package main

import (
	"fmt"
	"io/ioutil"
	"math"
	_ "oauth/auth"
	"oauth/config"
	_ "oauth/database"
	"oauth/database/bean"
	"oauth/database/iredis"
	"oauth/database/schema"
	"oauth/logger"
	"oauth/utils"
	"strings"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
)

var (
	cookieNameForSessionID = "mgtv-oauth-sessionid"
	session                = sessions.New(sessions.Config{
		Cookie:  cookieNameForSessionID,
		Expires: 45 * time.Minute, // <=0 means unlimited life
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

func GetApp() *iris.Application {
	app := iris.New()

	//conf := config.Get()
	app.Post("/authorize", func(ctx iris.Context) {
		token := ctx.GetHeader("token")
		clienID := ctx.GetHeader("client_id")
		account := ctx.GetHeader("account")

		if result, err := iredis.Get(fmt.Sprintf("%s:%s", clienID, account)); err != nil {
			ctx.StatusCode(403)
			return
		} else if result != token {
			ctx.StatusCode(403)
			return
		}

		body, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			logger.Debug(err)
			ctx.StatusCode(500)
			return
		}
		defer ctx.Request().Body.Close()

		app, err := bean.FindApplicationByClientID(clienID)

		if err != nil {
			logger.Debug(err)
			ctx.StatusCode(500)
			return
		}

		if app.ID == 0 {
			ctx.StatusCode(403)
			return
		}

		encryptKey := app.PrivateKey

		plainText, err := utils.CFBDecrypt(encryptKey, string(body))

		//
		ctx.StatusCode(200)
	})
	app.Get("/oauth", func(ctx iris.Context) {
		clientID := ctx.URLParam("client_id")
		now := time.Now().Unix()
		if timestamp, err := ctx.URLParamInt64("timestamp"); err != nil {
			ctx.StatusCode(406)
			return
			//与服务器误差小于5分钟 东八区
		} else if math.Abs(float64(now*1000-timestamp)) > 1000*5*60 {
			ctx.StatusCode(406)
			return
		}

		if clientID == "" {
			ctx.StatusCode(406)
			return
		}
		//是否存在应用ID
		app, err := bean.FindApplicationByClientID(clientID)

		if err != nil {
			ctx.StatusCode(500)
			ctx.WriteString(err.Error())
			return
		}

		if app.ID == 0 {
			ctx.StatusCode(406)
			return
		}

		sess := session.Start(ctx)
		//用户是否已登陆
		if userAuthorized, err := sess.GetBoolean("user-authorized"); err != nil || !userAuthorized {
			ctx.ServeFile("static/login.html", false)
			return
		}

		username := sess.Get("username")
		if username == "" {
			ctx.ServeFile("static/login.html", false)
			return
		}

		key := fmt.Sprintf("%s:%s", clientID, username)
		//使用应用私有Key进行 给 key+时间戳 进行加密，然后计算md5值
		encryptKey := utils.CFBEncrypt(app.PrivateKey, fmt.Sprintf("%s:%d", key, now))
		token := utils.MD5Str(encryptKey)

		if err := iredis.SetEx(key, token, 24*60*60*time.Second); err != nil {
			ctx.StatusCode(500)
			ctx.WriteString(err.Error())
			return
		}

		callback := fmt.Sprintf("%s?token=%s", app.Callback, token)

		ctx.StatusCode(301)
		ctx.Header("Location", callback)

	})

	app.Post("/oauth", func(ctx iris.Context) {
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
			ctx.JSON(app)
		}
	})

	return app
}

func main() {
	app := GetApp()
	app.Run(iris.Addr(":"+config.Get().Port), iris.WithoutVersionChecker)
}
