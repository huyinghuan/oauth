package main

import (
	"oauth/config"
	"oauth/controller"
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
	tmpl := iris.HTML("./static", ".html")
	tmpl.Reload(true)

	app.RegisterView(tmpl)

	app.Get("/", func(ctx iris.Context) {
		sess := session.Start(ctx)
		//用户是否已登陆
		if userAuthorized, err := sess.GetBoolean("user-authorized"); err != nil || !userAuthorized {
			ctx.ServeFile("static/login.html", false)
			return
		}
		ctx.ServeFile("static/user.html", false)
	})

	userCtrl := controller.User{Session: session}
	UserAPI := app.Party("/user")
	//注册
	UserAPI.Post("/register", userCtrl.Post)
	//退出
	UserAPI.Delete("/logout", userCtrl.Logout)
	//获取信息
	UserAPI.Get("/info", userCtrl.Get)
	//提交登陆表单
	UserAPI.Post("/login", userCtrl.Login)

	appCtrl := controller.App{Session: session}
	AppAPI := app.Party("/app")
	//注册
	AppAPI.Post("/register", appCtrl.Post)

	//以下为第三方调用接口
	authorizeCtrl := controller.Authorize{Session: session}
	app.Get("/authorize", authorizeCtrl.Get)
	//权限校验
	app.Post("/authorize", authorizeCtrl.Verity)

	//接口调整
	app.Post("/authorize/jump", authorizeCtrl.Jump)

	resourceCtrl := controller.Resource{Session: session}
	ResourceAPI := app.Party("/resource")
	ResourceAPI.Post("/account", resourceCtrl.GetAccount)

	return app
}

func main() {
	app := GetApp()
	app.Run(iris.Addr(":"+config.Get().Port), iris.WithoutVersionChecker)
}
