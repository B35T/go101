package main

import "fmt"

func main() {
	queue := make(chan int , 20)

	for i := 0; i < 10; i++ {
		queue <- i
	}
	close(queue) // ใช้งานเมื่อต้องการเพิ่มข้อมูล ใช้งานได้ครั้งเดียว

	for element := range queue {
		fmt.Println(element)
	}
}
