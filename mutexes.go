package main

import (
	"fmt"
	"time"
	"sync"
	"sync/atomic"
	"math/rand"
)

func main(){
	var state = make(map[int]int)
	// initialize an object of sync.Mutex
	fmt.Println("old state:", state)
	var mutex = &sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				fmt.Println("reader:", r, " access readOps is", readOps)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func(){
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				fmt.Println("writer:", w, " access writeOps is", writeOps)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	readOpsFinal := atomic.LoadUint64(&readOps)
	//readOpsFinal := readOps
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	//writeOpsFinal := writeOps
	fmt.Println("writeOps:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}