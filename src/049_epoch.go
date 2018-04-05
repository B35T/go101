package main

import (
	"time"
	"fmt"
)

func main() {
	now := time.Now()
	sec := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	milis := nanos / 1000000
	fmt.Println(sec)
	fmt.Println(nanos)
	fmt.Println(milis)

	fmt.Println(time.Unix(sec,0))
	fmt.Println(time.Unix(0, nanos))
}
