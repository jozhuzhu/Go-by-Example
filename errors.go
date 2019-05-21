package main

import "fmt"
import "errors"

func f1(argc int) (int, error){
	if argc == 42 {
		return -1, errors.New("can't work with 42")
	}
	return argc + 3, nil
}

type argError struct{
	arg int
	prob string
}

func (e *argError) Error() string{
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(argc int) (int, error){
	if argc == 42{
		return -1, &argError{argc, "can't work with it."}
	}
	return argc + 3, nil
}

func main(){
	for _, num := range []int{7, 42}{
		if r, e := f1(num); e != nil{
			fmt.Println("f1 failed:", e)
		}else{
			fmt.Println("f1 worked:", r)
		}
	}

	for _, num := range []int{7, 42}{
		if r, e := f2(num); e != nil{
			fmt.Println("f1 failed:", e)
		}else{
			fmt.Println("f1 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok{
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}