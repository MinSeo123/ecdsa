package base58Encode

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
)

func Base58Encode(ver []uint8 ,pubkeyHash []uint8, checkSum [] uint8 ) string {

	    data := append(ver[:], pubkeyHash[:]...)
	    data = append(data[:], checkSum[:]...)
	    fmt.Println(data)
		encoded := base58.Encode(data)

		return encoded
}