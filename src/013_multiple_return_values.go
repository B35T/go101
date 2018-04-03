package main

import "fmt"

type User struct {
	name string
	age int
	skill []string
}

func main() {
	users := User{
		name:"best",
		age:24,
		skill:[]string{"go","git"},
	}
	users.Print()

	fmt.Println(users.Value())
}

func (users User) Print() {
	fmt.Println(users)
}

func (users User) Value() (int,bool) {
	return 50, true
}
