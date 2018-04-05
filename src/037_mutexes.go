package main

import (
	"sync"
	"math/rand"
	"sync/atomic"
	"time"
	"fmt"
)

func main() {
	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	var readOperations uint64
	var writeOperations uint64

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOperations, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOperations, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOperationsFinal := atomic.LoadUint64(&readOperations)
	fmt.Println("read operation ", readOperationsFinal)

	writeOperationsFinal := atomic.LoadUint64(&writeOperations)
	fmt.Println("write operation ", writeOperationsFinal)

	mutex.Lock()
	fmt.Println("state :", state)
	mutex.Unlock()
}
