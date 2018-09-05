package auth

import (
	"encoding/json"
	"fmt"
	"oauth/database/iredis"
	"oauth/utils"
)

func DecryptBody(clientID string, body []byte) (string, error) {
	appPKKey := fmt.Sprintf("app:pk:%s", clientID)
	pk, err := iredis.Get(appPKKey)
	if err != nil {
		return "", err
	}
	return utils.CFBDecrypt(pk, string(body))
}

func EncryptBody(clientID string, data interface{}) (string, error) {
	body, _ := json.Marshal(data)
	appPKKey := fmt.Sprintf("app:pk:%s", clientID)
	pk, err := iredis.Get(appPKKey)
	if err != nil {
		return "", err
	}
	return utils.CFBEncrypt(pk, string(body))
}
