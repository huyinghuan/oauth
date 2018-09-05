package auth

import (
	"fmt"
	"oauth/database/iredis"
	"oauth/utils"
	"time"
)

func CreateResourceToken(clientID string, username string, pk string) (string, error) {
	encryptKey, err := utils.CFBEncrypt(pk, fmt.Sprintf("%s:%s:%d", clientID, username, time.Now().UnixNano()))
	if err != nil {
		return "", err
	}
	token := utils.MD5Str(encryptKey)
	key := fmt.Sprintf("resource:%s:%s", clientID, token)
	err = iredis.SetEx(key, username, 60*time.Second)
	return token, err
}

func GetResourceToken(clientID string, token string) (string, error) {
	key := fmt.Sprintf("resource:%s:%s", clientID, token)
	token, err := iredis.Get(key)
	if err != nil {
		return "", err
	}

	if err := iredis.Del(key); err != nil {
		return "", err
	}
	return token, nil
}
