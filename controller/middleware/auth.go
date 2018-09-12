package middleware

import (
	"github.com/kataras/iris/context"

	"github.com/kataras/iris/sessions"
)

type MiddleWare struct {
	Session *sessions.Sessions
}

func (m *MiddleWare) UserAuth(ctx context.Context) {
	sess := m.Session.Start(ctx)
	username := sess.GetString("username")
	if username == "" {
		ctx.StatusCode(401)
		return
	}
	ctx.Next()

}
