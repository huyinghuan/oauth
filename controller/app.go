package controller

import (
	"fmt"
	"log"
	"oauth/config"
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
	Callback string `json:"callback"`
}

//app修改页面

func (c *App) EditPage(ctx iris.Context) {
	appID, _ := ctx.Params().GetInt64("appID")
	sess := c.Session.Start(ctx)
	uid, err := sess.GetInt64("uid")
	if err != nil {
		ctx.StatusCode(302)
		ctx.Header("Location", "/")
		return
	}
	app, err := bean.GetAppliction(appID, uid)
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	if app.ID == 0 {
		ctx.StatusCode(404)
		return
	}
	ctx.ViewData("App", app)
	ctx.View("app-edit.html")
}

//app 修改

func (c *App) Put(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("appID")
	sess := c.Session.Start(ctx)
	uid, _ := sess.GetInt64("uid")
	form := appForm{}
	ctx.ReadJSON(&form)
	if err := bean.UpdateApplication(id, uid, &schema.Application{
		Name:     form.Name,
		Callback: form.Callback,
	}); err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	ctx.StatusCode(200)

}

// app 注册
func (c *App) Post(ctx iris.Context) {
	sess := c.Session.Start(ctx)
	username := sess.GetString("username")
	//如果没有开放应用注册或用户不是管理员，那么就拒绝注册
	if !config.Get().OpenRegister && username != config.Get().Account.User {
		ctx.StatusCode(401)
		return
	}
	if username == "" {
		ctx.StatusCode(401)
		return
	}
	form := appForm{}
	ctx.ReadJSON(&form)
	appName := strings.TrimSpace(form.Name)
	if appName == "" {
		ctx.StatusCode(iris.StatusNotAcceptable)
		ctx.WriteString("应用名不能为空")
		return
	}
	user, err := bean.FindUserByUsername(username)
	if err != nil {
		log.Println(err)
		ctx.StatusCode(401)
		return
	}
	app := schema.Application{
		UserID:   user.ID,
		Name:     appName,
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
