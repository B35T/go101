package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	//n,_ := os.Create("tmp/dat3")
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
		//t := []byte(ucl)
		//
		//w, _ := n.Write(t)
		//fmt.Println(string(w))
	}
	//defer n.Close()

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
