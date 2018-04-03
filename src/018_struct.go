package main

import "fmt"

type user struct {
	name string
	age int
}

func main() {
	fmt.Println(user{"best",23})

	fmt.Println(user{"bob",30})

	fmt.Println(user{"alice",32})

	fmt.Println(&user{"jabo",23})

	s := user{name:"pre",age:23}
	fmt.Println(s)

	fmt.Println(&s.name) // pointer address

	s.age = 26
	fmt.Println(s.age)
}
