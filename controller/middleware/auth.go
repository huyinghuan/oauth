package middleware

import (
	"time"

	"github.com/kataras/iris/v12/context"

	"github.com/kataras/iris/v12/sessions"
)


func UserAuth(ctx context.Context, session *sessions.Sessions) {
	session.UpdateExpiration(ctx, 30 * time.Minute)
	switch ctx.Path() {
	// 登陆请求跳过
	case "/api/user-status", "/api/user/register", "/api/open-register":
		ctx.Next()
	default:
		sess := sessions.Get(ctx)
		username := sess.GetString("username")
		if username == "" {
			ctx.StatusCode(401)
			ctx.WriteString("用户未登录")
			return
		}
		ctx.Next()
	}
}
