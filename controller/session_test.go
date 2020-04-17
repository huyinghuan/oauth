package controller

import (
  "testing"
  "time"

  "github.com/kataras/iris/v12/httptest"
)


/*
func TestMain(m *testing.M){
  // config.Init()
  // database.InitAll()
  sess := sessions.New(sessions.Config{
    Cookie:  cookieName,
    Expires: 30 * time.Minute,
  })

  app = iris.New()

  app.Use(sess.Handler())

  app.Get("/get", func(ctx iris.Context) {
    ctx.Writef("%v", sessions.Get(ctx).GetBooleanDefault("logged", false))
  })

  app.Get("/normal", func(ctx iris.Context) {
    sessions.Get(ctx).Set("logged", true)
    ctx.WriteString("OK")
  })


  app.Get("/cleanCookie", func(ctx iris.Context) {
    sessions.Get(ctx).Clear()
    sessions.Get(ctx).ClearFlashes()
    sessions.Get(ctx).Destroy()
    ctx.WriteString("OK")
  })

  app.Get("/remember_me", func(ctx iris.Context) {
    sessions.Get(ctx).Set("logged", true)
    sess.UpdateExpiration(ctx, 24*time.Hour)
    ctx.WriteString("OK")
  })

  app.Post("/post/remember_me", func(ctx iris.Context) {
    sessions.Get(ctx).Set("logged", true)
    sess.UpdateExpiration(ctx, 48*time.Hour)
    ctx.WriteString("OK")
  })
  exit := m.Run()
  os.Exit(exit)
}
*/
func TestFirstPostWillFail(t *testing.T){
  e := httptest.New(t, app, httptest.URL("http://example.com"))
  t.Run("Post Set expire", func(t *testing.T){
    tt := e.POST("/post/remember_me").Expect().Status(httptest.StatusOK)
    tt.Cookie(cookieName).MaxAge().Equal(48 * time.Hour - time.Second)
    tt.Body().Equal("OK")
    e.GET("/get").Expect().Status(httptest.StatusOK).Body().Equal("true")
  })
}

func TestSessionsUpdateExpiration(t *testing.T) {
  e := httptest.New(t, app, httptest.URL("http://example.com"))
  t.Run("normal", func(t *testing.T) {
    tt := e.GET("/normal").Expect().Status(httptest.StatusOK)
    tt.Cookie(cookieName).MaxAge().Equal(30 * time.Minute - time.Second)
    tt.Body().Equal("OK")
    // e.GET("/set").Expect().Status(httptest.StatusOK).Body().Equal("OK")
    e.GET("/get").Expect().Status(httptest.StatusOK).Body().Equal("true")
  })
  t.Run("clean", func(t *testing.T) {
    // clean
    e.GET("/cleanCookie").Expect().Status(httptest.StatusOK)
    e.GET("/get").Expect().Status(httptest.StatusOK).Body().Equal("false")
  })

  t.Run("Get Set expire", func(t *testing.T) {
    tt := e.GET("/remember_me").Expect().Status(httptest.StatusOK)
    tt.Cookie(cookieName).MaxAge().Equal(24 * time.Hour - time.Second)
    tt.Body().Equal("OK")
    e.GET("/get").Expect().Status(httptest.StatusOK).Body().Equal("true")
  })
  t.Run("clean", func(t *testing.T) {
    // clean
    e.GET("/cleanCookie").Expect().Status(httptest.StatusOK)
    e.GET("/get").Expect().Status(httptest.StatusOK).Body().Equal("false")
  })

  t.Run("Post Set expire", func(t *testing.T){
    tt := e.POST("/post/remember_me").Expect().Status(httptest.StatusOK)
    tt.Cookie(cookieName).MaxAge().Equal(48 * time.Hour - time.Second)
    tt.Body().Equal("OK")
    e.GET("/get").Expect().Status(httptest.StatusOK).Body().Equal("true")
  })
}