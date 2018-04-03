package main

import "fmt"

func main() {
	fmt.Println(fect(7))
}

func fect(n int) int {
	if n == 0 {
		return 1
	}
	return n * fect(n-1)
}
