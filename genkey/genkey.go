package genkey

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"os"
)

func GenKey()(pub []byte ,private *ecdsa.PrivateKey){
	//ECDSA KEYPAIR 생성
	pubkeyCurve := elliptic.P256()
	privateKey := new(ecdsa.PrivateKey)
	privateKey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var pubkey []byte
	pubkey = append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

	return pubkey, privateKey
}