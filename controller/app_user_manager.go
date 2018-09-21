package controller

import (
	"oauth/database/bean"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type AppUserManager struct {
	Session *sessions.Sessions
}

func (a *AppUserManager) GetView(ctx iris.Context) {
	appID, _ := ctx.Params().GetInt64("appID")
	app, err := bean.FindApplicationByID(appID)
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	ctx.ViewData("App", app)
	ctx.View("app-user.html")
}
