package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abcdef"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDnc,_ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDnc))

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	uDnc, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDnc))
}
