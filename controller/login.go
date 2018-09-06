package controller

import (
	"oauth/database/bean"
	"strings"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type Login struct {
	Session *sessions.Sessions
}
type accountForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Login) UserLogin(ctx iris.Context) {
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

func (c *Login) ApplicationLogin(ctx iris.Context) {

}
