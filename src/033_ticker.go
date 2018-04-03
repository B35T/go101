package main

import (
	"time"
	"fmt"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("ticker at:",t)
		}
	}()
	time.Sleep(2100 * time.Millisecond)
	ticker.Stop()
	fmt.Println("ticker stopped")
}
