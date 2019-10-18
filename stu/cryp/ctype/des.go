package ctype

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

//定义向量常量
const (
	iv   string = "12345678"         //8位向量
	iv16 string = "12345678abcdefgh" //16位向量
)

//补全明文的字节数组
func paddingDataText(plainText []byte, blockSize int) []byte {
	padNum := blockSize - len(plainText)%blockSize
	padByte := []byte{byte(padNum)}
	padBytes := bytes.Repeat(padByte, padNum)
	return append(plainText, padBytes...)

}

//去除明文的字节数组被补全的数据
func reducingDataText(plainText []byte) []byte {
	dataLen := len(plainText)
	lastByte := plainText[dataLen-1]
	num := int(lastByte)
	return plainText[:dataLen-num]
}

//CBC 方式加密数据
func DesEncryptByCBC(plainText, key []byte) (result []byte) {
	//创建并返回一个使用DES算法的cipher.Block接口，传入密钥
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Println("des new cripher error :", err)
	}
	//补全数据
	result = paddingDataText(plainText, block.BlockSize())
	//返回一个密码分组链接模式的、底层用b加密的BlockMode接口，初始向量iv的长度必须等于b的块尺寸。
	blockModle := cipher.NewCBCEncrypter(block, []byte(iv))
	//加密并返回结果
	blockModle.CryptBlocks(result, result)

	return

}

func DesDecryptByCBC(plainText, key []byte) (result []byte) {
	//创建并返回一个使用DES算法的cipher.Block接口，传入密钥
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Println("des new cripher error :", err)
	}
	//返回一个密码分组链接模式的、底层用b解密的BlockMode接口，初始向量iv必须和加密时使用的iv相同。
	blockModle := cipher.NewCBCDecrypter(block, []byte(iv))
	//解密
	blockModle.CryptBlocks(plainText, plainText)
	//移除被补全的数据
	result = reducingDataText(plainText)
	return

}
