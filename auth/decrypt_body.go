package auth

import (
	"encoding/json"
	"fmt"
	"oauth/database/iredis"

	SDK "github.com/huyinghuan/oauth_sdk"
)

func DecryptBody(clientID string, body []byte) (string, error) {
	appPKKey := fmt.Sprintf("app:pk:%s", clientID)
	pk, err := iredis.Get(appPKKey)
	if err != nil {
		return "", err
	}
	return SDK.CFBDecrypt(pk, string(body))
}

func EncryptBody(clientID string, data interface{}) (string, error) {
	body, _ := json.Marshal(data)
	appPKKey := fmt.Sprintf("app:pk:%s", clientID)
	pk, err := iredis.Get(appPKKey)
	if err != nil {
		return "", err
	}
	return SDK.CFBEncrypt(pk, string(body))
}
