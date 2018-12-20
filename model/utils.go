package model

import (
	"crypto/md5"
	"encoding/hex"
)

// GeneratePasswordHash 密码 md5 加密
func GeneratePasswordHash(pwd string) string {
	hasher := md5.New()
	hasher.Write([]byte(pwd))
	pwdHash := hex.EncodeToString(hasher.Sum(nil))
	return pwdHash
}
