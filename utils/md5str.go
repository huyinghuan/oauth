package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Str(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}
