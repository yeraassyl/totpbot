package main

import (
	"encoding/base32"
	"fmt"
)

func main(){
	secreteBase32 := "NNXXIYLLMJQXGZLCMF2CO3LBNRXGC2DV"
	secretBytes, _ := base32.StdEncoding.DecodeString(secreteBase32)
	fmt.Println(secretBytes)
}