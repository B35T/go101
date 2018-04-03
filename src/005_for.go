package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for x := 1;x <= 10 ;x++  {
		fmt.Println(x)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 10; n++ {
		if n%2 == 0 {
			var a = fmt.Sprint(n%2)
			fmt.Println("state = " + a)
			continue
		}
		fmt.Println(n)
	}
}
