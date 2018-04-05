package main

import (
	"os"
	"fmt"
)

func main() {
	argsWithProg := os.Args
	argsWhitoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWhitoutProg)
	fmt.Println(arg)
}
