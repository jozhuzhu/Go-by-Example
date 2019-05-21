package main

import "fmt"
import "time"

func main(){
	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		//for i := 0; i < 2; i++{
			time.Sleep(1 * time.Second)
			fmt.Println("enter the first function.")
			c1 <- "one"
		//}
	}()

	go func(){
		//for i := 0; i < 2; i++{
			time.Sleep(2 * time.Second)
			fmt.Println("enter the second function.")
			c2 <- "two"
		//}
	}()

	for i := 0; i < 2; i++ {
		select{
		case msg1 := <-c1:
			fmt.Println("received:", msg1)
		case msg2 := <-c2:
			fmt.Println("received:", msg2)
		}
	}
}