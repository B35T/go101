package main

import "fmt"

func main() {
	message := make(chan string, 2)

	message <- "golang"
	message <- "swift"

	fmt.Println(<-message)
	fmt.Println(<-message)
}
