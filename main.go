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


type Node struct {
	bitcoinAddress string
	signature []byte
	signhash []byte
	pubKey []byte
	privateKey *ecdsa.PrivateKey
}



func main() {
	var N Node
	//비트코인 주소 생성
	var createAddressNum int
	var signNum int
	var verifyNum int
	var data string
	fmt.Println("지갑을 생성하세요 1 입력")
	fmt.Println("사인하기 2 입력")
	fmt.Println("검증하기 3 입력")
	for true {
		fmt.Scanln(&createAddressNum)
		if createAddressNum == 1 {
			N.CreateAddress()
			fmt.Println("주소를 생성합니다, 사인하려면 2 입력", "주소:",N.bitcoinAddress)
			fmt.Scanln(&signNum)
			if signNum == 2 {
				fmt.Println("서명할 데이터를 입력하세요.")
				fmt.Scanln(&data)
				N.Signing(N.privateKey, data)
				fmt.Println("사인을 시작합니다, 검증하려면 3 입력 ","데이터 :", data )
				fmt.Scanln(&verifyNum)
				if verifyNum == 3{
					fmt.Scanln("검증을 시작합니다.")
					result := N.Verification(N.privateKey, N.signhash, N.signature)
					if result == true {
						fmt.Println("검증완료")
						break
					} else {
						fmt.Println("검증실패")
						break
					}
				} else if verifyNum != 3 {
					fmt.Println("명령어를 제대로 입력해주세요.")
					break
				}
			} else if signNum != 2 {
				fmt.Println("명령어를 제대로 입력해주세요. ")
				break
			}
		} else if createAddressNum == 2 || createAddressNum == 3{
			fmt.Println("지갑을 먼저 생성해주세요.")
		} else {
			fmt.Println("명령어를 제대로 입력하세요.")
		}
	}




	N.CreateAddress()

	N.Signing(N.privateKey, data)
	N.Verification(N.privateKey, N.signhash, N.signature)
}

func (N *Node) CreateAddress ()  string {
	//키페어 생성 (Pubkey, prikey)
	pubKey, privateKey := genkey.GenKey()
	N.pubKey = pubKey
	N.privateKey = privateKey
	//버전 설정
	version := []byte{1}
	//sha256을 이용한 해쉬 암호화
	shaResult := shaEncrypt.ShaEncrypt(N.pubKey)
	//RIPEMD
	ripemdResult := ripemdEncrypt.RipemdEncrypt(shaResult)
	PubkeyHash := []byte(ripemdResult)
	//Checksum
	checksumResult := shaEncrypt.ShaShaEncrypt(shaResult)
	CheckSum := checksumResult
	//baseEncode
	baseEncoded := base58Encode.Base58Encode(version, PubkeyHash, CheckSum)
	N.bitcoinAddress = baseEncoded
	return baseEncoded
}


func (N *Node) Signing (privateKey *ecdsa.PrivateKey, data string) ([]byte, []byte) {
	//사인
	signHash, signature := genkey.SignEcdsa(privateKey, data)
	N.signhash = signHash
	N.signature = signature
	return signHash, signature
}

func (N *Node) Verification(privatekey *ecdsa.PrivateKey, signhash []byte, signature []byte) bool {
	//검증
	verify := genkey.Verifycation(signhash, signature, privatekey)
	return verify
}