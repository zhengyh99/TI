package main

import (
	"cryp/ctype"
	"fmt"
)

func main() {
	str := "I love you ,LJX ! 我爱李晶先 !!!"
	// fmt.Println("des by CBC============================")
	// encode := ctype.DesEncryptByCBC([]byte(str), []byte("87654321"))
	// fmt.Println("密文:", string(encode))
	// decode := ctype.DesDecryptByCBC(encode, []byte("87654321"))
	// fmt.Println("明文:", string(decode))
	// fmt.Println("aes by CTR============================")
	// encode2 := ctype.AesComplieCode([]byte(str), []byte("87654321hgfedcba"))
	// fmt.Println("密文:", string(encode2))
	// decode2 := ctype.AesComplieCode(encode2, []byte("87654321hgfedcba"))
	// fmt.Println("明文:", string(decode2))

	// ctype.GenerateRsaKey("d:\\privKey.pem", "d:\\pubKey.pem", 1024)

	// cypherText := ctype.EncryptRsa([]byte(str), "d:\\pubKey.pem")
	// fmt.Println("cyperText:", cypherText)
	// plainText := ctype.DecryptRsa(cypherText, "d:\\privKey.pem")
	// fmt.Println("planText:", string(plainText))
	// key := "helloworld"
	// hmacText := ctype.GenerateHmac([]byte(str), []byte(key))
	// equal := ctype.VerifyHmac([]byte(str), []byte(key), hmacText)
	// fmt.Println("equl:", equal)

	// sign := ctype.SigntureRsa("d://privKey.pem", []byte(str))

	// ok := ctype.VerifySignRsa("d://pubKey.pem", []byte(str), sign)

	// fmt.Println(" verfySingRsa result is :", ok)

	// ctype.GenerateECDSAKey("d:\\ecdsaPubKey.pem", "d:\\ecdsaprivKey.pem", elliptic.P521())

	rCode, sCode := ctype.SigntureEcdsa("d:\\ecdsaprivKey.pem", []byte(str))
	fmt.Println("rCode:", rCode)
	fmt.Println("================")
	fmt.Println("sCode:", sCode)
	ecdsaOk := ctype.VerifyEcdsa("d:\\ecdsaPubKey.pem", []byte(str), rCode, sCode)
	fmt.Println("VerifyEcdsa result is ", ecdsaOk)

}
