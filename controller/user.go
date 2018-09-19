package controller

import (
	"oauth/config"
	"oauth/database/bean"
	"strings"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type User struct {
	Session *sessions.Sessions
}

//用户注册
func (c *User) Post(ctx iris.Context) {
	sess := c.Session.Start(ctx)
	//如果没有开放用户认证，用户不是管理员，那么就拒绝注册
	if !config.Get().OpenRegister && sess.GetString("username") != config.Get().Account.User {
		ctx.StatusCode(401)
		return
	}

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
}

func (c *User) Logout(ctx iris.Context) {
	sess := c.Session.Start(ctx)
	sess.Clear()
	sess.ClearFlashes()
	ctx.StatusCode(200)
}

type accountForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *User) Login(ctx iris.Context) {
	account := accountForm{}
	ctx.ReadJSON(&account)
	username := strings.TrimSpace(account.Username)
	password := strings.TrimSpace(account.Password)
	if username == "" || password == "" {
		ctx.StatusCode(403)
		return
	}

	if u, exist, err := bean.FindUser(username, password); err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
	} else if !exist {
		ctx.StatusCode(403)
	} else {
		sess := c.Session.Start(ctx)
		sess.Set("user-authorized", true)
		sess.Set("username", username)
		sess.Set("uid", u.ID)
		ctx.StatusCode(200)
	}
}
