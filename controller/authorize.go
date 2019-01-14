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
		ctx.WriteString("参数错误: clientID 不能为空")
		return
	}
	if account == "" {
		ctx.StatusCode(406)
		ctx.WriteString("参数错误: account 不能为空")
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
		msg := fmt.Sprintf("权限请求: %s:%s  状态:%d\n 用户不存在", clientID, account, 401)
		log.Println(msg)
		ctx.StatusCode(401)
		ctx.WriteString(msg)
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
		msg := fmt.Sprintf("权限请求: %s:%s  状态:%d\n 用户为黑名单", clientID, account, 403)
		log.Println(msg)
		ctx.StatusCode(403)
		ctx.WriteString(msg)
		return
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
		if !verifyAPIAccessPromission(list, scope.Name, scope.Type) {
			msg := fmt.Sprintf("权限请求 %s : %s : %s : %s 角色无权限:%d\n", clientID, account, scope.Name, scope.Type, 403)
			log.Println(msg)
			ctx.StatusCode(403)
			ctx.WriteString(msg)
			return
		}
	}
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
	msg := fmt.Sprintf("权限请求 %s : %s : %s : %s 授权成功:%d\n", clientID, account, scope.Name, scope.Type, 200)
	log.Println(msg)
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
	referer := ctx.GetHeader("referer")
	redirectURL := ctx.URLParam("redirect")
	if redirectURL == "" && strings.Index(referer, ctx.Host()) == -1 {
		redirectURL = referer
	}
	log.Println("redirect:", redirectURL)
	clientID := ctx.URLParam("client_id")
	clientID = strings.TrimSpace(clientID)
	if clientID == "" {
		ctx.StatusCode(406)
		ctx.WriteString("参数错误,client_id不能为空")
		return
	}
	//是否存在私有key
	if !iredis.AppCache.Exist(clientID) {
		ctx.StatusCode(406)
		ctx.WriteString("参数错误, client_id不存在")
		return
	}

	sess := c.Session.Start(ctx)
	//用户是否已登陆
	uid, err := sess.GetInt64("uid")
	if err != nil {
		//用户登陆
		ctx.Redirect(fmt.Sprintf("/authorize/login?client_id=%s&t=%s&redirect=%s", clientID, ctx.URLParam("t"), redirectURL))
		return
	}
	if uid == 0 {
		uid, _ = sess.GetInt64("adminID")
	}
	appID, err := iredis.AppCache.GetMap(clientID)
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}
	log.Println(appID)
	app, _ := bean.Application.Get(appID)
	if app.ID == 0 {
		ctx.StatusCode(406)
		ctx.WriteString("应用存不存在")
		return
	}
	//确认用户是否在正常访问名单
	if haveEnterPromise, err := bean.HaveEnterPromise(uid, appID); err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	} else if !haveEnterPromise {
		//没有访问权限
		ctx.ViewData("AppName", app.Name)
		ctx.View("no-promise.html")
		return
	}

	username := sess.GetString("username")
	//是否已经进过确认页面
	if agree, err := sess.GetBoolean(clientID); err != nil || !agree {
		ctx.ViewData("ClientID", clientID)
		ctx.ViewData("AppName", app.Name)
		ctx.ViewData("Redirect", redirectURL)
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
		ctx.WriteString("回调地址错误:" + e.Error())
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

func (c *Authorize) Login(ctx iris.Context) {
	ctx.ViewData("OpenRegister", config.Get().OpenRegister)
	ctx.View("login.html")
}
