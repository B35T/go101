package main

import "fmt"

func main() {
	fmt.Println("hello world")

	var x = "*"
	for i := 0; i <= 10; i++ {
		fmt.Println(x)
		x += "*"
	}
}
