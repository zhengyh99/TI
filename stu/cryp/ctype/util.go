package ctype

import (
	"crypto"
	"encoding/pem"
	"fmt"
	"os"
)

//hash 用特定的加密函数加密数据
func hashSum(hashType crypto.Hash, hashText []byte) (hashCode []byte) {
	myHash := hashType.New()
	myHash.Write(hashText)
	hashCode = myHash.Sum(nil)
	return
}

//PEM编码的结构block写入到本地磁盘文件
func writeFileByPem(fileName string, block pem.Block) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os create file:", fileName, "error : ", err)
	}
	//解码并写入文件
	pem.Encode(f, &block)
	f.Close()
}

//读取文件内容，返回字节切片
func readFile(fileName string) []byte {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("os open error :", err)
	}
	fInfo, err := f.Stat()
	if err != nil {
		fmt.Println("file status error:", err)
	}
	buf := make([]byte, fInfo.Size())
	_, err2 := f.Read(buf)
	if err2 != nil {
		fmt.Println("file read error:", err2)
	}
	f.Close()
	return buf
}
