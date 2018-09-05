package utils

import (
	"log"
	"testing"
)

func TestCFB(t *testing.T) {
	key := RandomString(24)
	plainText := RandomString(24) + ":" + "admin"

	log.Println("Plain text: ", plainText)

	encryptText, _ := CFBEncrypt(key, plainText)

	log.Println("encrypt Text: ", encryptText)

	decryptText, _ := CFBDecrypt(key, encryptText)

	if decryptText != plainText {
		log.Println("decrypt Text: ", decryptText)
		t.FailNow()
	}

	log.Println(MD5Str(encryptText))
}
