package controller

import (
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

func (c *App) GetList(ctx iris.Context) {
	sess := c.Session.Start(ctx)
	//用户是否已登陆
	uid, _ := sess.GetInt64("uid")
	if list, err := bean.GetApplictionList(uid); err != nil {
		ctx.StatusCode(500)
		ctx.Text(err.Error())
	} else {
		ctx.JSON(list)
	}
}

//app 获取
func (c *App) Get(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("appID")
	app, err := bean.Application.Get(id)
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	ctx.JSON(map[string]string{
		"name":     app.Name,
		"callback": app.Callback,
		"model":    app.Mode,
	})
}

//app 修改
func (c *App) Put(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("appID")
	sess := c.Session.Start(ctx)
	uid, _ := sess.GetInt64("uid")
	form := appForm{}
	ctx.ReadJSON(&form)
	app := schema.Application{
		Name:     form.Name,
		Callback: form.Callback,
	}
	if uApp, err := bean.UpdateApplication(id, uid, &app); err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())

	} else {
		iredis.AppCache.SetCallback(id, uApp.Callback)
		ctx.StatusCode(200)
	}

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
		if err := iredis.AppCache.SetAll(app.ID, app.PrivateKey, app.Callback, app.Mode); err != nil {
			ctx.StatusCode(500)
			ctx.WriteString(err.Error())
			return
		}
		if err := iredis.AppCache.SetMap(app.ID, app.ClientID); err != nil {
			ctx.StatusCode(500)
			ctx.WriteString(err.Error())
			return
		}
		ctx.JSON(map[string]string{
			"client_id":   app.ClientID,
			"private_key": app.PrivateKey,
		})
	}
}

//删除一个app
func (c *App) Delete(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("appID")
	if err := iredis.AppCache.Clear(id); err != nil {
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

//更新app的用户名单模式

func (c *App) UpdateRunMode(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("appID")
	mode := ctx.Params().Get("mode")

	if mode != "black" {
		mode = "white"
	}

	if err := bean.UpdateApplicationRunMode(id, mode); err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}

}
