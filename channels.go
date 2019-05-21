package main

import "fmt"

func main(){
	// messages := make(chan string)
	// go func(){
	// 	messages <- "ping"
	// }()

	// msg := <- messages
	// fmt.Println(msg)

	//var a []int
	var c1, c2, c3 chan int
	var i1, i2 int

	select{
	case i1 = <-c1:
		fmt.Println("received:", i1, " from c1\n")
	case c2 <- i2:
		fmt.Println("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):
		if ok{
			fmt.Println("received:", i3, " from c3")
		}else{
			fmt.Println("error happen")
		}
	default:
		print("stop")
	}
}

