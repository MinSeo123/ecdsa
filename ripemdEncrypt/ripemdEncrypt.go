package ripemdEncrypt

import (

	"golang.org/x/crypto/ripemd160"
)

func RipemdEncrypt(b []uint8 ) []uint8  {
	RIPE := ripemd160.New()
	RIPE.Write(b)
	RIPED := RIPE.Sum(b)
	return RIPED

}
