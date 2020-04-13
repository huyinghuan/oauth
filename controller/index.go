package controller

import (
	"log"
	"oauth/config"
	"oauth/database/bean"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type WebIndex struct {
	Session *sessions.Sessions
}

func (c *WebIndex) Get(ctx iris.Context) {
	sess := sessions.Get(ctx)
	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
	//用户是否已登陆
	uid, err := sess.GetInt64("uid")
	if err != nil {
		ctx.ViewData("OpenRegister", config.Get().OpenRegister)
		ctx.View("login.html")
		return
	}
	username := sess.GetString("username")
	ctx.ViewData("Account", username)
	//默认为普通用户
	isAdmin := uid == 0
	OpenAppRegister := config.Get().OpenAppRegister
	//是否开发用户注册
	ctx.ViewData("OpenAppRegister", OpenAppRegister)
	viewHTML := "user.html"
	//管理员
	if isAdmin {
		adminID, _ := sess.GetInt64("adminID")
		ctx.ViewData("AdminID", adminID)
		ctx.ViewData("OpenAppRegister", true)
		viewHTML = "admin.html"
	}

	if OpenAppRegister || isAdmin {
		list, err := bean.GetApplicationList(uid)
		if err != nil {
			log.Println(err)
			ctx.StatusCode(500)
			return
		}
		ctx.ViewData("AppList", list)
	}
	if isAdmin {
		userList, err := bean.GetAllUser()
		if err != nil {
			log.Println(err)
			ctx.StatusCode(500)
			return
		}
		ctx.ViewData("UserList", userList)
	}

	ctx.View(viewHTML)
}
