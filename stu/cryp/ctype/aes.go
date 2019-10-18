package ctype

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

//Aes ctr加密

func AesComplieCode(plainText, key []byte) []byte {
	//创建并返回一个使用DES算法的cipher.Block接口，传入密钥
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("aes newcipher error:", err)
	}
	//返回一个计数器模式的、底层采用block生成key流的Stream接口，初始向量iv的长度必须等于block的块尺寸。
	//stream接口代表一个流模式的加/解密器
	stream := cipher.NewCTR(block, []byte(iv16))
	//加密/解密
	stream.XORKeyStream(plainText, plainText)
	return plainText
}
