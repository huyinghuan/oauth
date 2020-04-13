package main

import (
  "testing"
  "time"

  "github.com/kataras/iris/v12"
  "github.com/kataras/iris/v12/context"
  "github.com/kataras/iris/v12/httptest"
  "github.com/kataras/iris/v12/sessions"
)

func TestSessionsUpdateExpiration(t *testing.T) {
  app := iris.New()

  cookieName := "mycustomsessionid"

  sess := sessions.New(sessions.Config{
    Cookie:  cookieName,
    Expires: 30 * time.Minute,
  })

  app.Use(sess.Handler())

  app.Get("/get", func(ctx context.Context) {
    ctx.Writef("%v", sessions.Get(ctx).GetBooleanDefault("logged", false))
  })

  app.Get("/set", func(ctx iris.Context) {
    sessions.Get(ctx).Set("logged", true)
    ctx.WriteString("OK")
  })

  app.Get("/remember_me", func(ctx iris.Context) {
    sess.UpdateExpiration(ctx, 24*time.Hour)
    ctx.WriteString("OK")
  })

  e := httptest.New(t, app, httptest.URL("http://example.com"))
  tt := e.GET("/get").Expect().Status(httptest.StatusOK)
  tt.Cookie(cookieName).MaxAge().Equal(30 * time.Minute)
  tt.Body().Equal("false")

  e.GET("/set").Expect().Status(httptest.StatusOK).Body().Equal("OK")
  e.GET("/get").Expect().Status(httptest.StatusOK).Body().Equal("true")

  tt = e.GET("/remember_me").Expect().Status(httptest.StatusOK)
  tt.Cookie(cookieName).MaxAge().Equal(24 * time.Hour)
  tt.Body().Equal("OK")
}