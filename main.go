package main

import (
	"oauth/config"
	"oauth/controller"
	"oauth/controller/middleware"
	_ "oauth/database"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

var (
	cookieNameForSessionID = "mgtv-oauth-sessionid"
	session                = sessions.New(sessions.Config{
		Cookie: cookieNameForSessionID,
		//Expires: 45 * time.Minute, // <=0 means unlimited life
	})
)

func GetApp() *iris.Application {
	app := iris.New()
	tmpl := iris.HTML("./static/template", ".html")
	tmpl.Reload(true)

	app.RegisterView(tmpl)

	app.StaticWeb("/static/", "./static/resource")

	//免登陆接口
	webIndexCtrl := controller.WebIndex{Session: session}
	app.Get("/", webIndexCtrl.Get)
	userCtrl := controller.User{Session: session}
	app.PartyFunc("/user", func(u iris.Party) {
		u.Get("/register", func(ctx iris.Context) { ctx.ServeFile("static/register.html", false) })
		//注册
		u.Post("/register", userCtrl.Post)
		//退出
		u.Delete("/logout", userCtrl.Logout)
		//提交登陆表单
		u.Post("/login", userCtrl.Login)
	})

	appCtrl := controller.App{Session: session}
	app.PartyFunc("/app", func(u iris.Party) {
		u.Get("/register", func(ctx iris.Context) { ctx.ServeFile("static/app-register.html", false) })
		//注册
		u.Post("/register", appCtrl.Post)
	})

	//需要登陆认证的接口
	middle := middleware.MiddleWare{Session: session}
	API := app.Party("/api", middle.UserAuth)
	API.PartyFunc("/app", func(u iris.Party) {
		u.Delete("/{appID:long}", appCtrl.Delete)
	})

	//以下为第三方调用接口
	authorizeCtrl := controller.Authorize{Session: session}
	app.PartyFunc("/authorize", func(u iris.Party) {
		u.Get("/", authorizeCtrl.Get)
		//权限校验
		u.Post("/", authorizeCtrl.Verity)
		//接口跳转
		u.Post("/jump", authorizeCtrl.Jump)
	})
	resourceCtrl := controller.Resource{Session: session}
	app.PartyFunc("/resource", func(u iris.Party) {
		u.Post("/account", resourceCtrl.GetAccount)
	})

	return app
}

func main() {
	app := GetApp()
	app.Run(iris.Addr(":"+config.Get().Port), iris.WithoutVersionChecker)
}
