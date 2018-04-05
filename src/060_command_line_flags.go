package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word","foo","a string")

	numPtr := flag.Int("numb", 42, "an int")
	boolP := flag.Bool("boolP",false,"a bool")

	var svar string
	flag.StringVar(&svar,"svar","bar","a string svar")
	flag.Parse()

	fmt.Println("word:",*wordPtr)
	fmt.Println("number:",*numPtr)
	fmt.Println("frok:", *boolP)
	fmt.Println("swar:", svar)
	fmt.Println("tail:",flag.Args())
}
