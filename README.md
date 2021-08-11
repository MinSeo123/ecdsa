![화면 기록 2021-08-11 오후 7 03 17](https://user-images.githubusercontent.com/80306757/129010874-03dfccb8-ccf0-4f4c-bfe4-ebc42acd0232.gif)
# 비트코인 주소 체계를 기반으로 임의의 데이터의 전자서명 검증 로직 개발 

## 실행방법
$go mod tidy<br/>
$go mod init<br/>
$go build<br/>

## 로직
ECDSA 기반 키 페어(public, private) 생성
비트코인의 주소 체계와 같은 주소 생성 (sha256, RIPEMD160)
인코드 (Base58Encode)

ECDSA(public key, private key) >  pubkeyHash > checksum > Base58Encode(version + pubkeyhash + checksum)

pubkeyHash = RIPEMD160(SHA256(Publickey)),
checksum = SHA256(SHA256(PubkeyHash))

privatekey와 데이터(signhash)를 더해 signature 생성 (signing)  
pubkey 와 signhash , signature를 더해 검증 (verify)
