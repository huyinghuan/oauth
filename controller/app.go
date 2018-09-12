package controller

import (
	"fmt"
	"log"
	"oauth/database/bean"
	"oauth/database/iredis"
	"oauth/database/schema"
	"strings"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type App struct {
	Session *sessions.Sessions
}
type appForm struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Callback string `json:"callback"`
}

// app 注册
func (c *App) Post(ctx iris.Context) {
	form := appForm{}
	ctx.ReadJSON(&form)
	username := strings.TrimSpace(form.Name)
	password := strings.TrimSpace(form.Password)
	if username == "" || password == "" {
		ctx.StatusCode(iris.StatusNotAcceptable)
		ctx.WriteString("用户名或密码不能为空")
		return
	}
	app := schema.Application{
		Name:     username,
		Password: password,
		Callback: form.Callback,
	}
	if err := bean.RegisterAppliction(&app); err != nil {
		ctx.StatusCode(iris.StatusNotAcceptable)
		ctx.WriteString(err.Error())
	} else {
		ctx.StatusCode(200)
		iredis.Set(fmt.Sprintf("app:pk:%s", app.ClientID), app.PrivateKey)
		iredis.Set(fmt.Sprintf("app:cb:%s", app.ClientID), app.Callback)
		ctx.JSON(map[string]string{
			"client_id":   app.ClientID,
			"private_key": app.PrivateKey,
		})
	}
}

//删除一个app
func (c *App) Delete(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("appID")
	app, err := bean.FindApplicationByID(id)
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}
	if err := iredis.Del(fmt.Sprintf("app:pk:%s", app.ClientID), fmt.Sprintf("app:cb:%s", app.ClientID)); err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}

	if err := bean.DeleteAppliction(id); err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}
	ctx.StatusCode(200)
}
