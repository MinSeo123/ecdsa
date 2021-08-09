package shaEncrypt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
)

func ShaEncrypt(a ecdsa.PublicKey) []uint8 {
	data := elliptic.Marshal(a, a.X, a.Y)
	hash := sha256.New()
	hash.Write([]byte(data))
	hashed := hash.Sum(data)
	return hashed
}

func ShaShaEncrypt(a []uint8) []uint8 {
	hash := sha256.New()
	hash.Write([]byte(a))
	hashed := hash.Sum(a)
	return hashed
}