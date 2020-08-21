package middleware

import (
	"log"
	"oauth/database/bean"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

func UserHaveApp(ctx iris.Context) {

	appID, _ := ctx.Params().GetInt64("appID")

	sess := sessions.Get(ctx)
	currentUID, _ := sess.GetInt64("uid")

	app, err := bean.Application.Get(appID)
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}
	// 如果app 不存在
	if app.ID == 0 {
		ctx.StatusCode(404)
		ctx.WriteString("找不到该应用")
		return
	}
	// 如果是管理员
	if currentUID == 0 {
		ctx.Next()
		return
	}
	if app.UserID != currentUID {
		ctx.StatusCode(403)
		return
	}

	ctx.Next()
}
