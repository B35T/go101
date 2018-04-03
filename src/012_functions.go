package main

import "fmt"

func main() {
	//name := User{
	//	name:"best",
	//	age:24,
	//	skill:[]string{"go","swift","git"},
	//}
	//
	//fmt.Println(name)


	a := plus(23,43)
	fmt.Println(a)
}

func plus(one int, two int) int {
	return  one + two
}
