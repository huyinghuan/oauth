package middleware

import (
	"github.com/kataras/iris/context"

	"github.com/kataras/iris/sessions"
)

type MiddleWare struct {
	Session *sessions.Sessions
}

func (m *MiddleWare) UserAuth(ctx context.Context) {
	switch ctx.Path() {
	//登陆请求跳过
	case "/api/user/login":
		ctx.Next()
		break
	default:
		sess := m.Session.Start(ctx)
		username := sess.GetString("username")
		if username == "" {
			ctx.StatusCode(401)
			return
		}
		ctx.Next()
	}

}
