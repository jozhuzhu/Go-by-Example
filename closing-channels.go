package main

import "fmt"
import "time"
func main(){
	jobs := make(chan int, 2)
	done := make(chan bool)

	go func(){
		i := 1
		for {
			fmt.Println("enter worker function:", i)
			i++
			time.Sleep(10 * time.Second)
			j, more := <- jobs
			if more{
				fmt.Println("received job:", j)
			}else{
				fmt.Println("no job received")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job:", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}