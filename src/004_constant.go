package main

import (
	"fmt"
	"math"
)

const name = "best"
func main() {
	fmt.Println(name)

	const pi  = math.Pi
	fmt.Println(pi)

	const n = 500000000
	fmt.Println(math.Sin(n))
}
