package main

import "fmt"

func main(){
	messages := make(chan string)
	signals := make(chan string)

	select{
	case message := <-messages:
		fmt.Println("received message:", message)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select{
	case messages <- msg:
		fmt.Println("sent message:", msg)
	default:
		fmt.Println("no message sent")
	}

	select{
	case msg := <-messages:
		fmt.Println("received message:", msg)
	case signal := <-signals:
		fmt.Println("received signal:", signal)
	default:
		fmt.Println("no activity")
	}
}