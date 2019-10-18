package ctype

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

const (
	//RSASSA-PKCS1-V1_5-SIGN签名方案计算签名 中的hashed算法
	hashType = crypto.SHA512
)

//生成RSA密码对，分别写入本地磁盘文件
func GenerateRsaKey(pvKeyFile, pubKeyFile string, bits int) {
	//生成私钥
	priv, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Println("rsa generate key error:", err)
	}
	//私钥解码
	privByte := x509.MarshalPKCS1PrivateKey(priv)
	//初始化私钥PEM编码的结构block
	block := pem.Block{
		Type:  " rsa private key ",
		Bytes: privByte,
	}
	//将私钥写入本地磁盘文件
	writeFileByPem(pvKeyFile, block)
	//从私钥中获取公钥
	pubKey := priv.PublicKey
	//公钥解码
	pubByte, err := x509.MarshalPKIXPublicKey(&pubKey)
	if err != nil {
		fmt.Println("x 509 MarshalPKIXPublicKey error:", err)
	}
	//初始化公钥PEM编码的结构block
	block = pem.Block{
		Type:  " rsa public key ",
		Bytes: pubByte,
	}
	//将公钥写入本地磁盘文件
	writeFileByPem(pubKeyFile, block)
}

//读取公钥文件 ，返回RSA公钥
func getRsaPubKeyFromFile(pubKeyFile string) (pubKey *rsa.PublicKey) {
	//读取公钥文件
	readPubKey := readFile(pubKeyFile)
	//解码 返回 pem block
	block, _ := pem.Decode(readPubKey)
	//解析公钥数据
	pubKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("x509 ParsePKIXPublicKey error:", err)
	}
	//断言，判断返回信息的type 是否为*rsa.PublicKey
	pubKey, ok := pubKeyInterface.(*rsa.PublicKey)
	if !ok {
		//不是*rsa.PublicKey 抛出异常
		fmt.Println("rsa public key interface error!")
		panic("rsa public key interface errorr")
	}
	//返回正确的 *rsa.PublicKey
	return
}

//读取私钥文件 ，返回RSA私钥
func getRsaPrivKeyFromFile(privKeyFile string) (privKey *rsa.PrivateKey) {
	//读取私钥文件
	readPrivKey := readFile(privKeyFile)
	//解码 返回pem block
	block, _ := pem.Decode(readPrivKey)
	//解析私钥数据
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("x509 ParsePKCS1PrivateKey error:", err)
	}
	//返回正确的 *rsa.PrivateKey
	return
}

//rsa公钥加密数据
func EncryptRsa(plainText []byte, pubKeyFile string) (cipherText []byte) {
	pubKey := getRsaPubKeyFromFile(pubKeyFile)
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, plainText)
	if err != nil {
		fmt.Println("rsa.EncryptPKCS1v15 error:", err)

	}
	return
}

//rsa私钥解密数据
func DecryptRsa(cipherText []byte, privKeyFile string) (plainText []byte) {
	privKey := getRsaPrivKeyFromFile(privKeyFile)
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, cipherText)
	if err != nil {
		fmt.Println("rsa.DecryptPKCS1v15 error:", err)
	}
	return
}

//生成 RSASSA-PKCS1-V1_5-SIGN签名方案计算签名
func SigntureRsa(privKeyFile string, plainText []byte) (sign []byte) {
	privKey := getRsaPrivKeyFromFile(privKeyFile)
	//hashType 常量 用特定的加密函数加密数据
	hashCode := hashSum(hashType, plainText)
	//SignPKCS1v15使用RSA PKCS#1 v1.5规定的RSASSA-PKCS1-V1_5-SIGN签名方案计算签名。
	//注意hashed必须是使用提供给本函数的hash参数对（要签名的）原始数据进行hash的结果。
	sign, err := rsa.SignPKCS1v15(rand.Reader, privKey, hashType, hashCode)
	if err != nil {
		fmt.Println("rsa sing pkcs1v15 error : ", err)
	}
	return
}

//认证RSA PKCS#1 v1.5签名
func VerifySignRsa(pubKeyFile string, plainText, signText []byte) bool {
	pubKey := getRsaPubKeyFromFile(pubKeyFile)
	//hashType 常量 用特定的加密函数加密数据
	hashCode := hashSum(hashType, plainText)
	//VerifyPKCS1v15认证RSA PKCS#1 v1.5签名。
	//hashed是使用提供的hash参数对（要签名的）原始数据进行hash的结果。
	//合法的签名会返回nil，否则表示签名不合法。
	err := rsa.VerifyPKCS1v15(pubKey, hashType, hashCode, signText)
	if err != nil {
		fmt.Println("rsa verifypkcs1v15 error ：", err)
		return false
	}
	return true

}
