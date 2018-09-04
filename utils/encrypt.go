package utils

import (
	"crypto/sha1"
	"fmt"
)

const salt = "docker-auth"

func Encrypt(content string) string {
	content = content + salt
	h := sha1.New()
	h.Write([]byte(content))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
