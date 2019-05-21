package main

import "fmt"

type rect struct{
	height, width int
}

func (r *rect) area(){
	r.width = 11
}

func (r rect) perim(){
	r.height = 77
}

func main(){
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.width)
	fmt.Println("perim:", r.height)

	rp := &r
	fmt.Println("area:", rp.width)
	fmt.Println("perim:", rp.height)
}