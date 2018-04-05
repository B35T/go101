package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main() {
	sign := make(chan os.Signal, 1)
	done := make(chan bool,1)

	signal.Notify(sign,syscall.SIGINT,syscall.SIGTERM)

	go func() {
		sig := <-sign
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("a waiting signal")
	<-done
	fmt.Println("exited")
}
