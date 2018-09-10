package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"oauth/auth"
	"oauth/database/bean"
	"oauth/database/iredis"
	"oauth/logger"
	"strings"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type Authorize struct {
	Session *sessions.Sessions
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

//权限校验

func (c *Authorize) Verity(ctx iris.Context) {
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
		Timestamp: time.Now().Unix(),
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
	if userAuthorized, err := sess.GetBoolean("user-authorized"); err != nil || !userAuthorized {
		ctx.StatusCode(401)
		return
	}
	sess.Set(clientID, true)
	ctx.StatusCode(200)
}

func (c *Authorize) Get(ctx iris.Context) {
	clientID := ctx.URLParam("client_id")
	clientID = strings.TrimSpace(clientID)
	if clientID == "" {
		ctx.StatusCode(406)
		return
	}

	//是否存在私有key
	appPKKey := fmt.Sprintf("app:pk:%s", clientID)

	if !iredis.Exist(appPKKey) {
		ctx.StatusCode(406)
		return
	}

	sess := c.Session.Start(ctx)
	//用户是否已登陆
	if userAuthorized, err := sess.GetBoolean("user-authorized"); err != nil || !userAuthorized {
		ctx.ServeFile("static/login.html", false)
		return
	}

	username := sess.GetString("username")
	if username == "" {
		ctx.ServeFile("static/login.html", false)
		return
	}
	if agree, err := sess.GetBoolean(clientID); err != nil || !agree {
		sess.Set(clientID, false)
		app, _ := bean.FindApplicationByClientID(clientID)
		ctx.ViewData("ClientID", clientID)
		ctx.ViewData("AppName", app.Name)
		ctx.View("confirm.html")
		return
	}

	privateKey, err := iredis.Get(appPKKey)
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
	cbURL, err := iredis.Get(fmt.Sprintf("app:cb:%s", clientID))
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	callback := fmt.Sprintf("%s?token=%s", cbURL, token)
	ctx.StatusCode(301)
	ctx.Header("Location", callback)
}
