package controller

import (
	"encoding/base64"
	"log"
	"oauth/auth"
	"oauth/config"
	"oauth/database/bean"
	"oauth/logger"
	"strings"

	"github.com/kataras/iris/context"
)

func verityUsernameAndPassword(account string, authorization string) bool {
	authorization = strings.Replace(authorization, "Basic ", "", -1)
	dencoded, err := base64.StdEncoding.DecodeString(authorization)
	if err != nil {
		logger.Debug(err)
		return false
	}
	arr := strings.Split(string(dencoded), ":")
	if len(arr) != 2 {
		return false
	}
	username, password := arr[0], arr[1]
	if account != username {
		logger.Debug("Error: Form Value 'account' %s != Header: 'Authorization' %s \n", account, username)
		return false
	}
	if _, exist, err := bean.FindUser(account, password); err != nil {
		log.Println(err)
		return false
	} else {
		return exist
	}
}

// TODO
func userInGroup(account string, group string) bool {
	return false
}

func isRepositoryOwner(account string, repositoryStr string) bool {
	arr := strings.Split(repositoryStr, "/")
	if len(arr) < 2 {
		return false
	}
	owner := arr[0]
	if account == owner {
		return true
	}
	return userInGroup(account, owner)
}

func verityUserScopePermission(form *auth.UserPostForm, hasLogin bool) {
	// TODO :  Database verity
	for index, scope := range form.Scopes {
		logger.Debugf("need vertiy scope => type: %s, name: %s, action: %s \n", scope.Type, scope.Name, scope.Actions)
		//admin account has login
		if hasLogin && form.Account == config.Get().Account.User {
			form.Scopes[index].Actions = scope.VerityActions
			continue
		}
		if scope.Type == "repository" {
			//if the account own this repository
			if hasLogin && isRepositoryOwner(form.Account, scope.Name) {
				form.Scopes[index].Actions = scope.VerityActions
				continue
			}
			form.Scopes[index].Actions = []string{"pull"}
		}

		if scope.Type == "registry" {
			form.Scopes[index].Actions = scope.VerityActions
		}
	}

}

func Authorize(ctx context.Context) {
	//token := ctx.GetHeader("token")

	// hasLogin := verityUsernameAndPassword(ctx.FormValue("account"), ctx.GetHeader("Authorization"))
	// form, err := auth.GetUserPostForm(ctx)
	// if err != nil {
	// 	logger.Debug(err)
	// 	ctx.StatusCode(iris.StatusBadRequest)
	// 	ctx.WriteString(fmt.Sprintf("Bad request: %s", err))
	// 	return
	// }

	// verityUserScopePermission(form, hasLogin)

	// tokenStr, err := auth.CreateToken(form)
	// if err != nil {
	// 	logger.Debug(err)
	// 	ctx.StatusCode(iris.StatusServiceUnavailable)
	// 	ctx.WriteString(err.Error())
	// 	return
	// }
	// result, _ := json.Marshal(&map[string]string{"token": tokenStr})
	// ctx.StatusCode(200)
	// ctx.Header("Content-Type", "application/json")
	// ctx.Write(result)
}
