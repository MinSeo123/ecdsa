package main

import (
	_ "crypto/sha256"
	"fmt"
	"github.com/MinSeo123/ecdsa/base58Encode"
	"github.com/MinSeo123/ecdsa/genkey"
	"github.com/MinSeo123/ecdsa/ripemdEncrypt"
	"github.com/MinSeo123/ecdsa/shaEncrypt"
	_ "golang.org/x/crypto/ripemd160"
)


type BitcoinAddress struct {
	Version []byte
	PubkeyHash []byte
	CheckSum []byte
}


func main() {
	var b BitcoinAddress

	//키페어 생성 (Pubkey, prikey)
	pubKey, _ := genkey.GenKey()
	b.Version = []byte{1}
	fmt.Println("펍키:",  pubKey)
	//sha256을 이용한 해쉬 암호화
	shaResult := shaEncrypt.ShaEncrypt(pubKey)
	fmt.Println("Sha256" , shaResult)
	fmt.Printf("sha256: %x", shaResult)

	//RIPEMD
	ripemdResult := ripemdEncrypt.RipemdEncrypt(shaResult)
	fmt.Println()
	fmt.Println("Ripemd :",ripemdResult)
	b.PubkeyHash = []byte(ripemdResult)


	//Checksum
	checksumResult := shaEncrypt.ShaShaEncrypt(shaResult)
	fmt.Println("체크섬" , checksumResult)
	fmt.Printf("체크섬: %x",checksumResult)
	b.CheckSum = checksumResult
	//b.CheckSum = fmt.Sprintf("%x", checksumResult)
	//fmt.Println("b.체크섬" , b.CheckSum)
	//baseEncode
	baseEncoded := base58Encode.Base58Encode(b.Version, b.PubkeyHash, b.CheckSum)
	fmt.Println("베이스인코드:", baseEncoded)
	//fmt.Println((baseEncoded))
	fmt.Println(BitcoinAddress{b.Version, b.PubkeyHash,b.CheckSum })





}