package middleware

import (
	"fmt"
	"log"
	"net/http"
	"oauth/config"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
)

type JWTUserInfo struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

func jwtGetUserInfo(tokenStr string) (info JWTUserInfo, code int, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.Get().JWTSecret), nil
	})
	if err != nil {
		return info, http.StatusUnauthorized, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return info, http.StatusUnauthorized, fmt.Errorf("验证不通过")
	}

	username := ""
	if v, exists := claims["username"]; exists {
		if n, ok := v.(string); ok {
			username = n
		}
	}
	info = JWTUserInfo{
		Username: username,
	}

	return info, 0, nil
}

func inIgnorePath(path string, ignorePaths []string) bool {
	for _, ignore := range ignorePaths {
		if path == ignore {
			return true
		}
	}
	return false
}

func FindUserFromJWT(mustLogin bool, ignorePaths []string) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		tokenStr := ctx.GetCookie("token")
		// 如果cookie中不存在　token则从url参数中读取
		if tokenStr == "" {
			tokenStr = ctx.URLParam("token")
		}
		// 如果还是为空，则登录用户为空
		if tokenStr == "" {
			if mustLogin && !inIgnorePath(ctx.Path(), ignorePaths) {
				ctx.JSON(Msg{
					Code: 401,
					Msg:  "需要登录",
				})
				return
			}
			ctx.Values().Set("jwt", JWTUserInfo{})
			ctx.Next()
			return
		}

		jwtInfo, statusCode, err := jwtGetUserInfo(tokenStr)
		if err != nil || statusCode != 0 {
			log.Println(err, statusCode)
			ctx.JSON(Msg{
				Code: 401,
				Msg:  "获取登录凭证失败，请重新登录",
			})
			return
		}
		ctx.Values().Set("jwt", jwtInfo)
		ctx.Next()
	}
}
