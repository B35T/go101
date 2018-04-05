package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 hash message"
	s2 := "sha1 hash message"
	salt := "qwerty"

	h := sha1.New()
	h.Write([]byte(s))
	bs1 := h.Sum(nil) // no add salt

	h2 := sha1.New()
	h2.Write([]byte(s2))
	bs2 := h2.Sum([]byte(salt)) // add salt

	fmt.Println(s)
	fmt.Println(bs1)
	fmt.Printf("%x\n", bs1)

	hash := fmt.Sprintf("%x",bs2)
	fmt.Println(hash)
}
