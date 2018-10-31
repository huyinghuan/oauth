package controller

import (
	"oauth/database/bean"
	"strings"

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
	list, _ := bean.Role.GetRoleList(appID)
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
	roleList, err := bean.Role.GetRoleList(appID)
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	ctx.ViewData("RoleList", roleList)
	if whiteList, err := bean.GetAppUserList(appID, "white"); err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	} else {
		ctx.ViewData("WhiteList", whiteList)
	}
	if blackList, err := bean.GetAppUserList(appID, "black"); err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	} else {
		ctx.ViewData("BlackList", blackList)
	}

	ctx.ViewData("App", app)

	ctx.View("app-user.html")
}

type appUserPostForm struct {
	Username string `json:"username"`
	Category string `json:"category"`
	RoleID   int64  `json:"role_id"`
}

func (a *AppUserManager) Post(ctx iris.Context) {
	appID, _ := ctx.Params().GetInt64("appID")
	form := appUserPostForm{}
	ctx.ReadJSON(&form)

	username := strings.TrimSpace(form.Username)

	category := strings.TrimSpace(form.Category)

	if username == "" || category == "" {
		ctx.StatusCode(406)
		return
	}
	user, err := bean.FindUserByUsername(username)
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	bean.AddUserToApp(user.ID, appID, category, form.RoleID)

}

func (a *AppUserManager) Delete(ctx iris.Context) {
	id, _ := ctx.Params().GetInt64("id")
	err := bean.DeleteUserFromAppUserList(id)
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
	}
}
