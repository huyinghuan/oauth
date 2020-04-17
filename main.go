package main

import (
	"oauth/config"
	"oauth/controller"
	"oauth/controller/middleware"
	"oauth/database"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/sessions"
)

var (
	cookieNameForSessionID = "mgtv-oauth-sessionid"
	session               = sessions.New(sessions.Config{
		Cookie:  cookieNameForSessionID,
		Expires:  0, // <=0 means unlimited life
		AllowReclaim: true,
	})
)

func GetApp() *iris.Application {
	app := iris.New()
	tmpl := iris.HTML("./static/template", ".html")
	tmpl.Reload(true)

	app.RegisterView(tmpl)
	app.HandleDir("/static/", "./static/resource")

	//免登陆接口
	// webIndexCtrl := controller.WebIndex{Session: session}
	// app.Get("/", webIndexCtrl.Get)
	app.Get("/", func(ctx iris.Context) { ctx.View("index.html") })
	//用户管理

	//以下为第三方调用接口
	app.PartyFunc("/authorize", func(u iris.Party) {
		u.Get("/", controller.AuthorizeCtrl.Get)
		//权限校验
		u.Post("/", controller.AuthorizeCtrl.Verify)
		//接口跳转
		u.Post("/jump", controller.AuthorizeCtrl.Jump)

		u.Get("/login", controller.AuthorizeCtrl.Login)
	})
	app.PartyFunc("/resource", func(u iris.Party) {
		u.Post("/account", controller.ResourceCtrl.GetAccount)
	})

	//数据接口
	API := app.Party("/api", session.Handler(), func(context iris.Context) {
		middleware.UserAuth(context, session)
	})

	API.Get("/open-register",  controller.UserCtrl.IsOpenRegister)

	API.PartyFunc("/user-status", func(u router.Party) {
		u.Get("/", controller.UserStatusCtrl.Get)
		u.Post("/", func(ctx iris.Context) {controller.UserStatusCtrl.Post(ctx, session)})
		u.Delete("/", controller.UserStatusCtrl.Delete)
	})

	API.PartyFunc("/user", func(u iris.Party) {
		u.Get("/info/{id:long}", controller.UserCtrl.GetAnyOneInfo)
		u.Get("/", controller.UserCtrl.GetList)
		u.Delete("/{uid:long}", controller.UserCtrl.DeleteUser)

		//注册
		u.Post("/register", controller.UserCtrl.Post)
		u.Put("/password", controller.UserCtrl.ResetPassword)
		u.Put("/{uid:long}/password", controller.UserCtrl.ResetPassword4Admin)

	})

	//获取所有注册的app列表
	API.PartyFunc("/open-apps", func(u iris.Party) {
		u.Get("/", controller.OpenAppsCtrl.GetList)
	})

	API.PartyFunc("/app", func(u iris.Party) {
		u.Get("/", controller.AppCtrl.GetList)
		u.Post("/register", controller.AppCtrl.Post)
	})
	application := API.Party("/app/{appID:long}", middleware.UserHaveApp)
	application.PartyFunc("/", func(u iris.Party) {
		u.Get("/", controller.AppCtrl.Get)
		u.Delete("/", controller.AppCtrl.Delete)
		u.Put("/", controller.AppCtrl.Put)
		u.Patch("/user_mode/{mode:string}", controller.AppCtrl.UpdateRunMode)
	})

	//黑白名单
	application.PartyFunc("/user", func(u iris.Party) {
		u.Get("/", controller.AppUserManagerCtrl.Get)
		u.Post("/", controller.AppUserManagerCtrl.Post)
		u.Delete("/{id: long}", controller.AppUserManagerCtrl.Delete)
		u.Get("/{id: long}/info", controller.AppUserManagerCtrl.GetUserInfo)
		u.Put("/{id:long}/role", controller.AppUserManagerCtrl.UpdateUserRole)
	})

	//应用角色
	roleCtrl := controller.AppRoleManager{Session: session}
	application.PartyFunc("/role", func(u iris.Party) {
		u.Post("/", roleCtrl.Post)
		u.Get("/", roleCtrl.Get)
		u.Delete("/{id:long}", roleCtrl.Delete)
		u.Get("/{id:long}", roleCtrl.GetOne)
	})

	permissionCtrl := controller.AppRolePermission{Session: session}
	RolePermission := application.Party("/role/{roleID:long}/permission", middleware.AppHaveRole)
	RolePermission.Get("/", permissionCtrl.Get)
	RolePermission.Post("/", permissionCtrl.Post)
	RolePermission.Delete("/{id:long}", permissionCtrl.Delete)

	return app
}

func main() {
	config.Init()
	database.InitAll()
	app := GetApp()
	app.Run(iris.Addr(":" + config.Get().Port))
}
