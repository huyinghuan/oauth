package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"oauth/auth"
	"oauth/database/iredis"
	"oauth/logger"
	"strings"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

type Resource struct {
	Session *sessions.Sessions
}
type resourceAccountForm struct {
	Timestamp int64  `json:"timestamp"`
	Token     string `json:"token"`
}

func (c *Resource) GetAccount(ctx iris.Context) {
	clientID := ctx.GetHeader("client_id")
	clientID = strings.TrimSpace(clientID)
	if clientID == "" {
		ctx.StatusCode(406)
		return
	}

	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		logger.Debug(err)
		ctx.StatusCode(500)
		return
	}
	defer ctx.Request().Body.Close()
	appID, err := iredis.AppCache.GetMap(clientID)
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		return
	}
	postStr, err := auth.DecryptBody(appID, body)

	if err != nil {
		logger.Debug(err)
		ctx.StatusCode(406)
		return
	}

	postData := resourceAccountForm{}

	if err := json.Unmarshal([]byte(postStr), &postData); err != nil {
		logger.Debug(err)
		ctx.StatusCode(500)
		return
	}

	if math.Abs(float64(time.Now().Unix()-postData.Timestamp)) > 5*60 {
		logger.Debug("时间戳超时")
		ctx.StatusCode(406)
		return
	}
	token := postData.Token
	username, err := auth.GetResourceByToken(clientID, token)
	if err != nil {
		ctx.StatusCode(406)
		return
	}
	result := map[string]interface{}{
		"timestamp": time.Now().Unix(),
		"username":  username,
	}
	encryptBody, err := auth.EncryptBody(appID, result)
	if err != nil {
		logger.Debug(err)
		ctx.StatusCode(500)
		return
	}
	ctx.WriteString(encryptBody)
}
