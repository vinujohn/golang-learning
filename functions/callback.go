package main

import "fmt"

func main(){
	fmt.Printf("Out of inner function\n")

	myFunc := func(x int) int{
		fmt.Printf("Printing from callback %d\n", x)
		return 2
	}

	testCallback(myFunc)
}

func testCallback(myFunc func(int) int){
	result := myFunc(4)
	fmt.Printf("Result of callback %d\n", result)
}
