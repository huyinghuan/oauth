package auth

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kataras/iris/context"
)

type UserPostForm struct {
	ClientID string
	Service  string
	Account  string
	Scopes   []Scope
}

type Scope struct {
	Type          string
	Name          string
	Actions       []string
	VerityActions []string
}

func parseScope(scopeArr []string) ([]Scope, error) {
	scopes := []Scope{}
	for _, scopeStr := range scopeArr {
		parts := strings.Split(scopeStr, ":")
		var scope Scope
		switch len(parts) {
		case 3:
			scope = Scope{
				Type:          parts[0],
				Name:          parts[1],
				Actions:       []string{},
				VerityActions: strings.Split(parts[2], ","),
			}
		case 4:
			scope = Scope{
				Type:          parts[0],
				Name:          parts[1] + ":" + parts[2],
				Actions:       []string{},
				VerityActions: strings.Split(parts[3], ","),
			}
		default:
			return nil, fmt.Errorf("invalid scope: %q", scopeStr)
		}
		sort.Strings(scope.Actions)
		scopes = append(scopes, scope)
	}
	return scopes, nil
}

func GetUserPostForm(ctx context.Context) (form *UserPostForm, err error) {
	form = &UserPostForm{}
	form.ClientID = ctx.FormValue("client_id")
	form.Service = ctx.FormValue("service")
	form.Account = ctx.FormValue("account")
	scope := ctx.FormValues()["scope"]
	if len(scope) > 0 {
		form.Scopes, err = parseScope(scope)
		if err != nil {
			return
		}
	}
	return
}
