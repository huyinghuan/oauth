package middleware

import (
	"oauth/database/bean"

	"github.com/kataras/iris"
)

func (m *MiddleWare) UserHaveApp(ctx iris.Context) {
	sess := m.Session.Start(ctx)

	currentUID, _ := sess.GetInt64("uid")

	//如果是管理员
	if currentUID == -1 {
		ctx.Next()
		return
	}

	appID, _ := ctx.Params().GetInt64("appID")

	app, err := bean.FindApplicationByID(appID)
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
