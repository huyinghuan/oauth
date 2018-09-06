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

	authorizeCtrl := controller.Authorize{Session: session}
	app.Post("/authorize", authorizeCtrl.Post)
	app.Get("/authorize", authorizeCtrl.Get)

	userCtrl := controller.User{Session: session}
	app.Post("/user-register", userCtrl.Post)

	appCtrl := controller.App{Session: session}
	app.Post("/app-register", appCtrl.Post)

	resourceCtrl := controller.Resource{Session: session}
	ResourceAPI := app.Party("/resource")
	ResourceAPI.Post("/account", resourceCtrl.GetAccount)

	loginCtrl := controller.Login{Session: session}
	LoginAPI := app.Party("/login")
	LoginAPI.Post("/user", loginCtrl.UserLogin)
	LoginAPI.Post("/app", loginCtrl.ApplicationLogin)
	return app
}

func main() {
	app := GetApp()
	app.Run(iris.Addr(":"+config.Get().Port), iris.WithoutVersionChecker)
}
