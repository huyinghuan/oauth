package auth

import (
	"encoding/json"
	"oauth/database/iredis"

	SDK "github.com/huyinghuan/oauth_sdk"
)

func DecryptBody(clientID string, body []byte) (string, error) {
	pk, err := iredis.AppCache.GetPrivateKey(clientID)
	if err != nil {
		return "", err
	}
	return SDK.CFBDecrypt(pk, string(body))
}

func EncryptBody(clientID string, data interface{}) (string, error) {
	body, _ := json.Marshal(data)
	pk, err := iredis.AppCache.GetPrivateKey(clientID)
	if err != nil {
		return "", err
	}
	return SDK.CFBEncrypt(pk, string(body))
}
