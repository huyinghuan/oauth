package controller

import (
	"log"
	"oauth/config"
	"oauth/database/bean"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type WebIndex struct {
	Session *sessions.Sessions
}

func (c *WebIndex) Get(ctx iris.Context) {
	sess := c.Session.Start(ctx)
	//用户是否已登陆
	if userAuthorized, err := sess.GetBoolean("user-authorized"); err != nil || !userAuthorized {
		ctx.ServeFile("static/login.html", false)
		return
	}
	username := sess.GetString("username")
	if username != config.Get().Account.User {
		ctx.ServeFile("static/user.html", false)
		return
	}
	ctx.ViewData("Account", username)

	list, err := bean.GetApplictionList()
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}
	ctx.ViewData("AppList", list)

	userList, err := bean.GetAllUser()

	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}
	ctx.ViewData("UserList", userList)

	ctx.View("admin.html")
}
