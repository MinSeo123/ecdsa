package ripemdEncrypt

import (
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

func RipemdEncrypt(b []byte ) string  {
	RIPE := ripemd160.New()
	RIPE.Write(b)
	RIPED := RIPE.Sum(nil)
	RIPEDHASH := fmt.Sprintf("%x", RIPED)
	return RIPEDHASH

}
