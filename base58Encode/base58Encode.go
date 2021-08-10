package base58Encode

import (
	"github.com/btcsuite/btcutil/base58"
)

func Base58Encode(ver []byte ,pubkeyHash []byte, checkSum []byte ) string {

		data := append(ver, pubkeyHash... )
		data2 := append(data, checkSum...)
	    //data := []byte(ver + pubkeyHash + checkSum)
		address := base58.Encode(data2)
		//encoded := base58.Encode(data)
	    //data := append(ver[:], pubkeyHash[:]...)
	    //data = append(data[:], checkSum[:]...)
		//encoded := base58.Encode(data)
		return address
}