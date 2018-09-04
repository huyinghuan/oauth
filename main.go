package main

import (
	_ "oauth/auth"
	"oauth/config"
	"oauth/controller"
	_ "oauth/database"
	"oauth/database/bean"
	"strings"

	"github.com/kataras/iris/context"

	"github.com/kataras/iris"
)

func GetApp() *iris.Application {
	app := iris.New()
	conf := config.Get()
	app.Any("/authorize", controller.Authorize)
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
