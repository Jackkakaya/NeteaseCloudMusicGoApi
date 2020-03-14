package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"math/big"
)

func RsaEncryptRaw(plainText []byte, publicKey []byte)  []byte {
	block, _ := pem.Decode(publicKey)
	pubInterface,err := x509.ParsePKIXPublicKey(block.Bytes)
	if err!=nil{
		panic(err)
	}
	pubKey := pubInterface.(*rsa.PublicKey)
	encrypted := new(big.Int)
	e := big.NewInt(int64(pubKey.E))
	payload := new(big.Int).SetBytes(plainText)
	encrypted.Exp(payload, e, pubKey.N)
	return encrypted.Bytes()
}