package ctype

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
)

func GenerateECDSAKey(pubKeyFile, privKeyFile string, curser elliptic.Curve) {
	privKey, err := ecdsa.GenerateKey(curser, rand.Reader)
	if err != nil {
		fmt.Println("ecdsa generatekey error:", err)
	}
	privKeyBytes, err := x509.MarshalECPrivateKey(privKey)
	if err != nil {
		fmt.Println("x509 marshal ecprivatekey error :", err)
	}
	block := pem.Block{
		Type:  " ecdsa private key ",
		Bytes: privKeyBytes,
	}
	writeFileByPem(privKeyFile, block)

	pubKey := privKey.PublicKey
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(&pubKey)
	if err != nil {
		fmt.Println("x509.MarshalPKIXPublicKey error:", err)
	}
	block = pem.Block{
		Type:  " ecdsa public key ",
		Bytes: pubKeyBytes,
	}
	writeFileByPem(pubKeyFile, block)
}

func getEcdsaPubKeyFromFile(fileName string) (pubKey *ecdsa.PublicKey) {
	readPubKey := readFile(fileName)
	block, _ := pem.Decode(readPubKey)
	pubKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("x509.ParsePKIXPublicKey error:", err)
	}
	pubKey, ok := pubKeyInterface.(*ecdsa.PublicKey)
	if !ok {
		//不是*ecdsa.PublicKey 抛出异常
		fmt.Println("ecdsa public key interface error!")
		panic("ecdsa public key interface errorr")
	}
	return
}

func getEcdsaPrivKeyFromFile(fileName string) (privKey *ecdsa.PrivateKey) {
	readPubKey := readFile(fileName)
	block, _ := pem.Decode(readPubKey)
	privKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("x509.ParseECPrivateKey error", err)
	}
	return
}

func SigntureEcdsa(privKeyFile string, plainText []byte) (rCode, sCode []byte) {
	privKey := getEcdsaPrivKeyFromFile(privKeyFile)
	hash := hashSum(crypto.SHA512, plainText)
	r, s, err := ecdsa.Sign(rand.Reader, privKey, hash)
	if err != nil {
		fmt.Println("ecdsa sing error :", err)
	}
	rCode, err = r.MarshalText()
	if err != nil {
		fmt.Println("r marshalText error:", err)
	}
	sCode, err = s.MarshalText()
	if err != nil {
		fmt.Println("s marshalText error:", err)
	}
	return
}

func VerifyEcdsa(pubKeyFile string, plainText, rCode, sCode []byte) bool {
	pubKey := getEcdsaPubKeyFromFile(pubKeyFile)
	hash := hashSum(crypto.SHA512, plainText)
	var r, s big.Int
	r.UnmarshalText(rCode)
	s.UnmarshalText(sCode)
	return ecdsa.Verify(pubKey, hash, &r, &s)
}
