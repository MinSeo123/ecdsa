package genkey

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"hash"
	"io"
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

func SignEcdsa(privateKey *ecdsa.PrivateKey, data string) ([]byte, []byte) {
	var h hash.Hash
	err := errors.New("can't sign")
	h = md5.New()
	var signature []byte
	//r := big.NewInt(0)
	//s := big.NewInt(0)
	io.WriteString(h, data)
	signhash := h.Sum(nil)

	signature, serr := ecdsa.SignASN1(rand.Reader, privateKey, signhash)
	if serr != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//signature := r.Bytes()
	//signature = append(signature, s.Bytes()...)
	//KeySign := fmt.Sprintf("%x\n", signature)
	//Verify
	//verifystatus := ecdsa.Verify(&pubkey, signhash, r, s)
	return signhash ,signature
}

func Verifycation(signhash []byte, signature []byte, privateKey *ecdsa.PrivateKey ) bool {
	var pubkey ecdsa.PublicKey
	pubkey = privateKey.PublicKey
	verifystatus := ecdsa.VerifyASN1(&pubkey, signhash, signature)
	return verifystatus

}