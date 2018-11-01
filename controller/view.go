package controller

import (
	"oauth/database/bean"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type View struct {
	Session *sessions.Sessions
}

func (a *View) GetAppUserView(ctx iris.Context) {
	appID, _ := ctx.Params().GetInt64("appID")
	app, err := bean.Application.Get(appID)
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	roleList, err := bean.Role.GetRoleList(appID)
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	ctx.ViewData("RoleList", roleList)

	ctx.ViewData("App", app)

	ctx.View("app-user.html")
}

func (a *View) GetAppRoleView(ctx iris.Context) {
	appID, _ := ctx.Params().GetInt64("appID")
	app, _ := bean.Application.Get(appID)
	list, _ := bean.Role.GetRoleList(appID)
	ctx.ViewData("RoleList", list)
	ctx.ViewData("App", app)
	ctx.View("app-user-role.html")
}

func (a *View) GetRolePermissionView(ctx iris.Context) {
	appID, _ := ctx.Params().GetInt64("appID")
	app, _ := bean.Application.Get(appID)
	roleID, _ := ctx.Params().GetInt64("roleID")
	role, _ := bean.Role.Get(roleID)
	permission, _ := bean.Role.GetPermission(roleID)
	ctx.ViewData("Role", role)
	ctx.ViewData("App", app)
	ctx.ViewData("PermissionList", permission)
	ctx.View("app-user-role-permission.html")
}

//app修改页面

func (c *View) GetAppEditPage(ctx iris.Context) {
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
