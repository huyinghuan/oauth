package controller

import (
	"oauth/database/bean"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type AppUserManager struct {
	Session *sessions.Sessions
}

func (a *AppUserManager) GetRolePermissionView(ctx iris.Context) {
	appID, _ := ctx.Params().GetInt64("appID")
	app, _ := bean.FindApplicationByID(appID)
	roleID, _ := ctx.Params().GetInt64("roleID")
	role, _ := bean.Role.Get(roleID)
	permission, _ := bean.Role.GetPermission(roleID)
	ctx.ViewData("Role", role)
	ctx.ViewData("App", app)
	ctx.ViewData("PermissionList", permission)
	ctx.View("app-user-role-permission.html")
}

func (a *AppUserManager) GetRoleView(ctx iris.Context) {
	appID, _ := ctx.Params().GetInt64("appID")
	app, _ := bean.FindApplicationByID(appID)
	list, _ := bean.Role.GetRoleList(app.ClientID)
	ctx.ViewData("RoleList", list)
	ctx.ViewData("App", app)
	ctx.View("app-user-role.html")
}

func (a *AppUserManager) GetView(ctx iris.Context) {
	appID, _ := ctx.Params().GetInt64("appID")
	app, err := bean.FindApplicationByID(appID)
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	if whiteList, err := bean.GetAppUserList(app.ClientID, "white"); err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	} else {
		ctx.ViewData("WhiteList", whiteList)
	}
	if blackList, err := bean.GetAppUserList(app.ClientID, "black"); err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	} else {
		ctx.ViewData("BlackList", blackList)
	}

	ctx.ViewData("App", app)

	ctx.View("app-user.html")
}

func (a *AppUserManager) Post(ctx iris.Context) {
	appID, _ := ctx.Params().GetInt64("appID")
	app, err := bean.FindApplicationByID(appID)

	form := map[string]string{}

	ctx.ReadJSON(&form)

	username, isExistUsername := form["username"]
	category, isExistCategory := form["category"]

	if !isExistUsername || !isExistCategory {
		ctx.StatusCode(406)
		return
	}
	user, err := bean.FindUserByUsername(username)
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}

	bean.AddUserToApp(user.ID, app.ClientID, category)

}

func (a *AppUserManager) Delete(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("id")
	err := bean.DeleteUserFromAppUserList(id)
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
	}
}
