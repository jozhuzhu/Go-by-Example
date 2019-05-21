package main

import "fmt"
import "sort"

func main(){
	var strs = []string{"c", "a", "d"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, -1}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}