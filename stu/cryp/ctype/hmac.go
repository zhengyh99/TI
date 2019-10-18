package ctype

import (
	"crypto/hmac"
	"crypto/sha1"
)
//HMAC（Hash-based Message Authentication Code）是密钥相关的哈希运算消息认证码

//生成HMAC 认证
func GenerateHmac(plainText, key []byte) []byte {
	mHash := hmac.New(sha1.New, key)
	mHash.Write(plainText)
	return mHash.Sum(nil)
}

//校验HMAC
func VerifyHmac(plainText, key, hmacText []byte) bool {
	hText := GenerateHmac(plainText, key)
	return hmac.Equal(hText, hmacText)
}
