package main

import "fmt"

func main() {

	var v = value(1,2,3,4,5)
	fmt.Println(v)

	var a = value(1)
	fmt.Println(a)
}

func value(num ...int) int {
	sum := 0
	for _,n := range num {
		sum += n
	}

	return sum
}
