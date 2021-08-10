package shaEncrypt

import (
	"crypto/sha256"
)

func ShaEncrypt(a []byte) []byte {
	hash := sha256.Sum256(a)
	//hash.Write([]byte(data))
	//hashed := hash.Sum(data)
	result := hash[0:len(hash)]
	return result
}

func ShaShaEncrypt(a []byte) []byte {
	hash := sha256.Sum256(a)
	//hash.Write([]byte(a))
	//hashed := hash.Sum(a)
	result := hash[0:len(hash)]
	return result
}