package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"oauth/auth"
	"oauth/config"
	"oauth/database/bean"
	"oauth/database/iredis"
	"oauth/database/schema"
	"regexp"
	"strconv"
	"strings"
	"time"

	SDK "github.com/huyinghuan/oauth_sdk"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type Authorize struct {
	Session *sessions.Sessions
}

// type Scope struct {
// 	Type    string
// 	Name    string
// 	Actions []string
// }

// type AuthScope struct {
// 	Timestamp int64 `json:"timestamp"`
// 	Scope     Scope `json:"scope"`
// }

func isAllowMethod(allowMethodList string, beTestMethod string) bool {
	methodList := strings.Split(allowMethodList, ",")
	//校验method
	for _, allowMethod := range methodList {
		allowMethod = strings.ToUpper(allowMethod)
		if allowMethod == "ALL" || allowMethod == beTestMethod {
			return true
		}
	}
	return false
}

func isAllowAPI(allowAPI string, url string) bool {
	if allowAPI == url {
		return true
	}
	reg, err := regexp.Compile(allowAPI)
	if err != nil {
		return false
	}
	return reg.MatchString(url)
}

/*
	method == GET,POST,PUT
*/
//API访问权限校验
func verifyAPIAccessPromission(list []schema.AppRolePermission, url string, method string) bool {
	for _, rule := range list {
		if !isAllowMethod(rule.Method, method) {
			continue
		}
		if !isAllowAPI(rule.Pattern, url) {
			continue
		}
		return true
	}
	return false
}

//权限校验
func (c *Authorize) Verify(ctx iris.Context) {
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
	//校验 account 是否存在数据库，是否处于正常状态
	user, err := bean.FindUserByUsername(account)
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}
	if user.ID == 0 {
		ctx.StatusCode(401)
		return
	}
	appID, err := iredis.AppCache.GetMap(clientID)
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}
	//是否设置了应用的黑白名单，当前用户是否拥有进入应用权限
	if haveEnterPromise, err := bean.HaveEnterPromise(user.ID, appID); err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	} else if !haveEnterPromise {
		ctx.StatusCode(401)
		return
	}
	//TODO 用户是否存在角色，该角色是否具有该路径的访问权限。
	roleID, err := bean.Role.GetRoleIDByUserIDInApp(appID, user.ID)
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}
	//角色为默认角色时，拥有全部权限
	if roleID != 0 {
		list, err := bean.Role.GetPermission(roleID)
		if err != nil {
			log.Println(err)
			ctx.StatusCode(500)
			return
		}
		if !verifyAPIAccessPromission(list, ctx.Path(), ctx.Method()) {
			ctx.StatusCode(403)
			return
		}
	}

	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}
	defer ctx.Request().Body.Close()
	decryptBody, err := auth.DecryptBody(appID, body)
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}
	scope := SDK.Scope{}
	if err := json.Unmarshal([]byte(decryptBody), &scope); err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}

	log.Println(fmt.Sprintf("权限请请求 %s : %s : %s : %s", clientID, account, scope.Name, scope.Type))

	authScope := SDK.AuthScope{
		Timestamp: time.Now().Unix(),
		Scope:     scope,
	}

	encryptAuthScope, err := auth.EncryptBody(appID, authScope)
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}
	ctx.StatusCode(200)
	ctx.WriteString(encryptAuthScope)
}

func (c *Authorize) Jump(ctx iris.Context) {
	clientID := ctx.URLParam("client_id")
	clientID = strings.TrimSpace(clientID)
	if clientID == "" {
		ctx.StatusCode(406)
		return
	}
	sess := c.Session.Start(ctx)
	//用户是否已登陆
	if _, err := sess.GetInt64("uid"); err != nil {
		ctx.StatusCode(401)
		return
	}
	sess.Set(clientID, true)
	ctx.StatusCode(200)
}

func (c *Authorize) Get(ctx iris.Context) {
	redirectURL := ctx.URLParam("redirect")
	clientID := ctx.URLParam("client_id")
	clientID = strings.TrimSpace(clientID)
	if clientID == "" {
		ctx.StatusCode(406)
		return
	}
	//是否存在私有key

	if !iredis.AppCache.Exist(clientID) {
		ctx.StatusCode(406)
		return
	}

	sess := c.Session.Start(ctx)
	//用户是否已登陆
	uid, err := sess.GetInt64("uid")
	if err != nil {
		ctx.ViewData("OpenRegister", config.Get().OpenRegister)
		ctx.View("login.html")
		return
	}
	appID, err := iredis.AppCache.GetMap(clientID)
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}
	//确认用户是否在正常访问名单
	if haveEnterPromise, err := bean.HaveEnterPromise(uid, appID); err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	} else if !haveEnterPromise {
		//没有访问权限
		app, _ := bean.FindApplicationByID(appID)
		ctx.ViewData("AppName", app.Name)
		ctx.View("no-promise.html")
		return
	}

	username := sess.GetString("username")
	//是否已经进过确认页面
	if agree, err := sess.GetBoolean(clientID); err != nil || !agree {
		app, _ := bean.FindApplicationByID(appID)
		ctx.ViewData("ClientID", clientID)
		ctx.ViewData("AppName", app.Name)
		ctx.View("confirm.html")
		return
	}
	//每次都需要进入确认页面
	sess.Set(clientID, false)
	privateKey, err := iredis.AppCache.GetPrivateKey(appID)
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	token, err := auth.CreateResourceToken(clientID, username, privateKey)
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	cbURL, err := iredis.AppCache.GetCallback(appID)
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	u, e := url.Parse(cbURL)
	if e != nil {
		ctx.StatusCode(406)
		return
	}
	q := u.Query()
	q.Add("redirect", redirectURL)
	q.Add("token", token)
	q.Add("t", strconv.FormatInt(time.Now().Unix(), 10))
	u.RawQuery = q.Encode()
	ctx.StatusCode(302)
	ctx.Header("Location", u.String())
}
