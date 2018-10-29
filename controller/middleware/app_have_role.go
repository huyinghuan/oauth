package middleware

import (
	"oauth/database/bean"

	"github.com/kataras/iris"
)

func AppHaveRole(ctx iris.Context) {
	appID, _ := ctx.Params().GetInt64("appID")
	app, err := bean.FindApplicationByID(appID)
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	id, _ := ctx.Params().GetInt64("id")
	exist, err := bean.Role.AppHaveRole(id, app.ClientID)
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	if exist {
		ctx.Next()
	} else {
		ctx.StatusCode(406)
		return
	}
}
