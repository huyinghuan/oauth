package controller

import (
	"oauth/database/bean"
	"oauth/database/schema"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type AppRolePermission struct {
	Session *sessions.Sessions
}

func (ctrl *AppRolePermission) Post(ctx iris.Context) {
	roleID, _ := ctx.Params().GetInt64("roleID")
	p := schema.AppRolePermission{}
	ctx.ReadJSON(&p)
	p.RoleID = roleID
	if err := bean.Perssion.Add(&p); err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	ctx.StatusCode(200)
}

func (ctrl *AppRolePermission) Delete(ctx iris.Context) {
	ID, _ := ctx.Params().GetInt64("id")
	if err := bean.Perssion.Remove(ID); err != nil {
		ctx.StatusCode(500)
		ctx.WriteString(err.Error())
		return
	}
	ctx.StatusCode(200)
}
