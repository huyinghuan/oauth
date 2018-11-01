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
