package main

import (
	"sort"
	"fmt"
)

func main() {
	strs := []string{"a","bing","go","zing","and","ant","man","full"}
	sort.Strings(strs)
	fmt.Println("string",strs)

	numbr := []int{3,1,5,3,87,45,2,9,0}
	sort.Ints(numbr)
	fmt.Println("number",numbr)

	s := sort.IntsAreSorted(numbr)
	fmt.Println(s)
}
