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

type appUserPostForm struct {
	Username string `json:"username"`
	Category string `json:"category"`
	RoleID   int64  `json:"role_id"`
}

func (a *AppUserManager) Get(ctx iris.Context) {
	appID, _ := ctx.Params().GetInt64("appID")
	list, err := bean.GetAppUserList(appID, "")
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	ctx.JSON(list)
}

func (a *AppUserManager) Post(ctx iris.Context) {
	appID, _ := ctx.Params().GetInt64("appID")
	form := appUserPostForm{}
	ctx.ReadJSON(&form)

	username := strings.TrimSpace(form.Username)

	category := strings.TrimSpace(form.Category)

	if username == "" || category == "" {
		ctx.StatusCode(406)
		ctx.Text("参数错误" + username + ":" + category)
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

func (a *AppUserManager) UpdateUserRole(ctx iris.Context) {
	form := map[string]int64{}
	if err := ctx.ReadJSON(&form); err != nil {
		ctx.StatusCode(406)
		ctx.WriteString("提交数据错误")
		return
	}
	roleID, ok := form["roleID"]
	if !ok {
		ctx.StatusCode(406)
		ctx.WriteString("提交数据错误")
		return
	}
	appID, _ := ctx.Params().GetInt64("appID")
	userID, _ := ctx.Params().GetInt64("id")
	status, err := bean.UpdateUserRoleInApp(appID, userID, roleID)
	if status != 200 {
		ctx.StatusCode(status)
		ctx.WriteString(err)
	}
	ctx.StatusCode(200)
}

func (a *AppUserManager) GetUserInfo(ctx iris.Context) {
	appID, _ := ctx.Params().GetInt64("appID")
	userID, _ := ctx.Params().GetInt64("id")
	user, err := bean.GetAppUserInfo(appID, userID)
	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
	}
	ctx.JSON(user)
}
