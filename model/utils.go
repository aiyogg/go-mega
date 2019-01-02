package model

import (
	"crypto/md5"
	"encoding/hex"
)

// GeneratePasswordHash 密码 md5 加密
func GeneratePasswordHash(pwd string) string {
	return Md5(pwd)
}

// Md5 func 使用 md5 加密字符串
func Md5(origin string) string {
	hasher := md5.New()
	hasher.Write([]byte(origin))
	return hex.EncodeToString(hasher.Sum(nil))
}
