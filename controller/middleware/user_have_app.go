package middleware

import (
	"oauth/database/bean"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

func UserHaveApp(ctx iris.Context) {

	appID, _ := ctx.Params().GetInt64("appID")

	sess := sessions.Get(ctx)
	currentUID, _ := sess.GetInt64("uid")
	//如果是管理员
	if currentUID == 0 {
		ctx.Next()
		return
	}

	app, err := bean.Application.Get(appID)
	if err != nil {
		ctx.StatusCode(500)
		return
	}
	if app.UserID != currentUID {
		ctx.StatusCode(403)
		return
	}
	ctx.Next()
}
