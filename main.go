package main

import (
	"crypto/ecdsa"
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

type Node struct {
	bitcoinAddress string
	signature []byte
	signhash []byte
	pubKey []byte
	privateKey *ecdsa.PrivateKey
}



func main() {
	var B BitcoinAddress
	var N Node
	data := "hi my name is minseo"
	//키페어 생성 (Pubkey, prikey)
	pubKey, privateKey := genkey.GenKey()
	//비트코인 주소 생성
	bitcoinAddress := B.CreateAddress(pubKey)
	N.pubKey = pubKey
	N.privateKey = privateKey
	N.bitcoinAddress = bitcoinAddress
	//사인
	signHash, signature := genkey.SignEcdsa(N.privateKey, data)
	N.signhash = signHash
	N.signature = signature
	//검증
	verify := genkey.Verifycation(N.signhash, N.signature, N.privateKey)
	fmt.Println(verify)
}

func (B *BitcoinAddress) CreateAddress (pubkey []byte)  string {
	B.Version = []byte{1}
	//sha256을 이용한 해쉬 암호화
	shaResult := shaEncrypt.ShaEncrypt(pubkey)
	//RIPEMD
	ripemdResult := ripemdEncrypt.RipemdEncrypt(shaResult)
	B.PubkeyHash = []byte(ripemdResult)
	//Checksum
	checksumResult := shaEncrypt.ShaShaEncrypt(shaResult)
	B.CheckSum = checksumResult
	//baseEncode
	baseEncoded := base58Encode.Base58Encode(B.Version, B.PubkeyHash, B.CheckSum)
	return baseEncoded
}