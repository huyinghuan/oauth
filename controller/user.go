package controller

import (
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

func (c *User) Get(ctx iris.Context) {
	sess := c.Session.Start(ctx)
	ctx.JSON(map[string]string{
		"username": sess.GetString("username"),
	})
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

	if _, exist, err := bean.FindUser(username, password); err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
	} else if !exist {
		ctx.StatusCode(403)
	} else {
		sess := c.Session.Start(ctx)
		sess.Set("user-authorized", true)
		sess.Set("username", username)
		ctx.StatusCode(200)
	}
}
