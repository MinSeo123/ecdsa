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
	"math/big"
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

func SignEcdsa(privateKey *ecdsa.PrivateKey) (string ,bool){
	var h hash.Hash
	err := errors.New("can't sign")
	h = md5.New()
	r := big.NewInt(0)
	s := big.NewInt(0)
	var pubkey ecdsa.PublicKey
	pubkey = privateKey.PublicKey
	io.WriteString(h, "This is a message to be signed and verified by ECDSA!")
	signhash := h.Sum(nil)

	r, s, serr := ecdsa.Sign(rand.Reader, privateKey, signhash)
	if serr != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	KeySign := fmt.Sprintf("Signature: %x\n", signature)
	fmt.Println(KeySign)

	//Verify
	verifystatus := ecdsa.Verify(&pubkey, signhash, r, s)
	fmt.Println(verifystatus)

	return KeySign, verifystatus
}