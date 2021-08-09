# 비트코인 주소 체계를 기반으로 임의의 데이터의 전자서명 검증 로직 개발 

ECDSA 기반 키 페어(public, private) 생성
비트코인의 주소 체계와 같은 주소 생성 (sha256, RIPEMD160)
인코드 (Base58Encode)

ECDSA(public key, private key) >  pubkeyHash > checksum > Base58Encode(version + pubkeyhash + checksum)

pubkeyHash = RIPEMD160(SHA256(Publickey))
checksum = SHA256(SHA256(PubkeyHash))
