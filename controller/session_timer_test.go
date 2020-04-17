package controller

import (
  "os"
  "sync"
  "testing"
  "time"

  "github.com/kataras/iris/v12"
  "github.com/kataras/iris/v12/httptest"
  "github.com/kataras/iris/v12/sessions"
)

var app *iris.Application
var cookieName = "mycustomsessionid"

func TestMain(m *testing.M){
  // config.Init()
  // database.InitAll()
  sess := sessions.New(sessions.Config{
    Cookie:  cookieName,
    Expires: 0,
    AllowReclaim: true,
  })

  app = iris.New()

  app.Use(sess.Handler())
  app.Use(func(ctx iris.Context){
    // session will expire after 30 minute at the last visit
    sess.UpdateExpiration(ctx, 30 * time.Minute)
    ctx.Next()
  })

  app.Get("/set", func(ctx iris.Context) {
    sessions.Get(ctx).Set("hello", "world")
    ctx.StatusCode(200)
  })

  app.Get("/get", func(ctx iris.Context) {
    ctx.WriteString(sessions.Get(ctx).GetString("hello"))
  })

  exit := m.Run()
  os.Exit(exit)
}

func TestSessionTimer(t *testing.T){
  e := httptest.New(t, app, httptest.URL("http://example.com"))

  e.GET("/set").Expect().Status(httptest.StatusOK)
  i := 0
  wg := sync.WaitGroup{}
  wg.Add(100)
  for i < 100 {
    go func() {
      tt := e.GET("/get").Expect().Status(httptest.StatusOK)
      tt.Body().Equal("world")
      tt.Cookie(cookieName).MaxAge().Equal(30 * time.Minute - time.Second)
      wg.Done()
    }()
    i++
  }
  wg.Wait()
  tt := e.GET("/get").Expect()
  tt.Status(httptest.StatusOK).Body().Equal("world")
  tt.Cookie(cookieName).MaxAge().Equal(30 * time.Minute - time.Second)
}
