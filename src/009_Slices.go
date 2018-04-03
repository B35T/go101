package main

import "fmt"

func main() {
	s := make([]string,3)
	fmt.Println("emp:",s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	fmt.Println("get:",s[2])

	fmt.Println("len:", len(s))

	s = append(s, "e")
	s = append(s, "s","f")
	fmt.Println("adp:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy :", c)

	l := c[2:6]
	fmt.Println(l)

	l = s[:5]
	fmt.Println(l)

	l = s[3:]
	fmt.Println(l)

	t := []string{"g","h","i"}
	fmt.Println(t)

	towD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		towD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			towD[i][j] = i + j
		}
	}
	fmt.Println(towD)
}
