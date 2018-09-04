package main

import (
	_ "oauth/auth"
	"oauth/config"
	"oauth/controller"
	_ "oauth/database"
	"oauth/database/bean"
	"strings"
	"time"

	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"

	"github.com/kataras/iris"
)

var (
	cookieNameForSessionID = "mgtv-oauth-sessionid"
	session                = sessions.New(sessions.Config{
		Cookie:  cookieNameForSessionID,
		Expires: 45 * time.Minute, // <=0 means unlimited life
	})
)

func GetApp() *iris.Application {
	app := iris.New()
	conf := config.Get()
	app.Any("/authorize", controller.Authorize)
	app.Get("/oauth", func(ctx iris.Context) {
		client_id := ctx.URLParam("client_id")
		if client_id == "" {
			ctx.StatusCode(406)
			return
		}

		app, err := bean.FindApplicationByClientID(client_id)

		if err != nil {
			ctx.StatusCode(500)
			ctx.WriteString(err.Error())
			return
		}

		if app.ID == 0 {
			ctx.StatusCode(406)
			return
		}

		sess := session.Start(ctx)

		if userAuthorized, err := sess.GetBoolean("user-authorized"); err != nil || !userAuthorized {
			//todo send login file

			return
		}

	})
	app.Post("/oauth", func(ctx iris.Context) {

	})
	app.Post("/"+conf.Account.API, func(ctx context.Context) {
		username := strings.TrimSpace(ctx.FormValue("username"))
		password := strings.TrimSpace(ctx.FormValue("password"))
		if username != "" && password != "" {
			if err := bean.RegisterUser(username, password); err != nil {
				ctx.StatusCode(iris.StatusNotAcceptable)
				ctx.WriteString(err.Error())
			} else {
				ctx.StatusCode(200)
				ctx.WriteString("注册成功！")
			}

		} else {
			ctx.StatusCode(iris.StatusNotAcceptable)
			ctx.WriteString("用户名或密码不能为空")
		}
	})
	return app
}

func main() {
	app := GetApp()
	app.Run(iris.Addr(":"+config.Get().Port), iris.WithoutVersionChecker)
}
