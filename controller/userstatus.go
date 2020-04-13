package controller

import (
  "net/http"
  "strings"
  "time"

  "oauth/config"
  "oauth/database/bean"

  "github.com/kataras/iris/v12"
  "github.com/kataras/iris/v12/sessions"
)

// 用户状态管理
type UserStatus struct {}

func (c *UserStatus) Delete(ctx iris.Context) {
  sess := sessions.Get(ctx)
  sess.Destroy()
  ctx.StatusCode(200)
}

type accountForm struct {
  Username string `json:"username"`
  Password string `json:"password"`
  Remember bool `json:"remember"`
}

func (c *UserStatus) Post(ctx iris.Context, sessionCtrl *sessions.Sessions) {
  account := accountForm{}
  if err := ctx.ReadJSON(&account); err!=nil{
    ctx.StatusCode(http.StatusNotAcceptable)
    ctx.WriteString("提交数据错误")
    return
  }
  username := strings.TrimSpace(account.Username)
  password := strings.TrimSpace(account.Password)
  if username == "" || password == "" {
    ctx.StatusCode(http.StatusNotAcceptable)
    ctx.WriteString("账户或密码不能为空")
    return
  }

  if u, exist, err := bean.FindUser(username, password); err != nil {
    ctx.StatusCode(500)
    ctx.WriteString(err.Error())
  } else if !exist {
    ctx.StatusCode(http.StatusNotAcceptable)
    ctx.WriteString("用户不存在")
  } else {
    sess := sessions.Get(ctx)
    sess.Set("username", username)
    //管理员
    if config.Get().Account.IsAdmin(username) {
      sess.Set("uid", 0)
      sess.Set("adminID", u.ID)
    } else {
      sess.Set("uid", u.ID)
    }
    if account.Remember {
      sessionCtrl.UpdateExpiration(ctx, 30 * 24 * time.Hour)
    }
    ctx.StatusCode(200)

  }
}


func (c *UserStatus) Get(ctx iris.Context) {
  sess := sessions.Get(ctx)
  uid, _ := sess.GetInt64("uid")
  username := sess.GetString("username")
  if username == "" {
    ctx.StatusCode(401)
    ctx.WriteString("用户未登录")
    return
  }
  ctx.JSON(map[string]interface{}{
    "username": username,
    "uid":      uid,
  })
}

var UserStatusCtrl UserStatus
