package main

import "time"
import "fmt"

func main(){
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)
	// limiter controls the speed of executing code
	// limiter simulate one channel with 1 capacility
	for req := range requests{
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	// burstyLimiter can make the first 3 requests faster due to its capacility of 3
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++{
		burstyLimiter <- time.Now()
	}

	go func(){
		for t := range time.Tick(200 * time.Millisecond){
			burstyLimiter <- t
		}
	}()

	burstyRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)

	for req := range burstyRequest{
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}