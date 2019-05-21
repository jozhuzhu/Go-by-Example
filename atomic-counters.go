package main

import "fmt"
import "time"
import "sync/atomic"

func main(){
	var ops uint64

	for i := 1; i <= 50; i++ {
		go func(){
			atomic.AddUint64(&ops, 1)

			time.Sleep(time.Millisecond)
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("opsFinal addresss;", &opsFinal)
	fmt.Println("ops address:", &ops)
	fmt.Println("ops:", opsFinal)
}