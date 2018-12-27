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
	// webIndexCtrl := controller.WebIndex{Session: session}
	// app.Get("/", webIndexCtrl.Get)
	app.Get("/", func(ctx iris.Context) { ctx.View("index.html") })
	//用户管理
	middle := middleware.MiddleWare{Session: session}

	//以下为第三方调用接口
	authorizeCtrl := controller.Authorize{Session: session}
	app.PartyFunc("/authorize", func(u iris.Party) {
		u.Get("/", authorizeCtrl.Get)
		//权限校验
		u.Post("/", authorizeCtrl.Verify)
		//接口跳转
		u.Post("/jump", authorizeCtrl.Jump)

		u.Get("/login", authorizeCtrl.Login)
	})
	resourceCtrl := controller.Resource{Session: session}
	app.PartyFunc("/resource", func(u iris.Party) {
		u.Post("/account", resourceCtrl.GetAccount)
	})

	//数据接口
	userCtrl := controller.User{Session: session}
	appCtrl := controller.App{Session: session}
	appUserMangerCtrl := controller.AppUserManager{Session: session}

	API := app.Party("/api", middle.UserAuth)
	API.PartyFunc("/user", func(u iris.Party) {
		u.Get("/", userCtrl.GetLoginUserInfo)
		u.Get("/info/{id:long}", userCtrl.GetAnyOneInfo)
		u.Get("/list", userCtrl.GetList)
		u.Post("/login", userCtrl.Login)
		u.Delete("/logout", userCtrl.Logout)
		//注册
		u.Post("/register", userCtrl.Post)
		u.Put("/password", userCtrl.ResetPassword)
		u.Put("/password/{uid:long}", userCtrl.ResetPassword4Admin)
		u.Delete("/{uid:long}", userCtrl.DeleteUser)
	})
	API.PartyFunc("/app", func(u iris.Party) {
		u.Get("/", appCtrl.GetList)
		u.Post("/register", appCtrl.Post)
	})
	application := API.Party("/app/{appID:long}", middle.UserHaveApp)
	application.PartyFunc("/", func(u iris.Party) {
		u.Get("/", appCtrl.Get)
		u.Delete("/", appCtrl.Delete)
		u.Put("/", appCtrl.Put)
		u.Patch("/user_mode/{mode:string}", appCtrl.UpdateRunMode)
	})

	//黑白名单
	application.PartyFunc("/user", func(u iris.Party) {
		u.Get("/", appUserMangerCtrl.Get)
		u.Post("/", appUserMangerCtrl.Post)
		u.Delete("/{id: long}", appUserMangerCtrl.Delete)
		u.Put("/{id:long}/role", appUserMangerCtrl.UpdateUserRole)
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
	RolePermission := application.Party("/role/{roleID:long}/permission", middle.AppHaveRole)
	RolePermission.Get("/", permissionCtrl.Get)
	RolePermission.Post("/", permissionCtrl.Post)
	RolePermission.Delete("/{id:long}", permissionCtrl.Delete)

	return app
}

func main() {
	app := GetApp()
	app.Run(iris.Addr(":"+config.Get().Port), iris.WithoutVersionChecker)
}
