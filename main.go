package main

import (
	_ "crypto/sha256"
	"fmt"
	"github.com/MinSeo123/ecdsa/base58Encode"
	"github.com/MinSeo123/ecdsa/genkey"
	"github.com/MinSeo123/ecdsa/ripemdEncrypt"
	"github.com/MinSeo123/ecdsa/shaEncrypt"
	_ "golang.org/x/crypto/ripemd160"
	"reflect"
)

func main() {
	//키페어 생성 (Pubkey, prikey)
	pubKey, _ := genkey.GenKey()
	ver := []uint8{1}

	//sha256을 이용한 해쉬 암호화
	shaResult := shaEncrypt.ShaEncrypt(pubKey)
	fmt.Println("Sha256" , shaResult)
	fmt.Printf("sha256: %x", shaResult)
	//RIPEMD
	ripemdResult := ripemdEncrypt.RipemdEncrypt(shaResult)

	fmt.Println("Ripemd",ripemdResult)
	fmt.Printf("Ripemd: %x", ripemdResult)

	//Checksum
	checksumResult := shaEncrypt.ShaShaEncrypt(shaResult)
	fmt.Println("체크섬" , checksumResult)
	fmt.Printf("체크섬: %x",checksumResult)

	//baseEncode
	baseEncoded := base58Encode.Base58Encode(ver ,shaResult,checksumResult)
	fmt.Println("베이스인코드:", baseEncoded)
	fmt.Println(reflect.TypeOf(baseEncoded))
}