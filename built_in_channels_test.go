package learning

import (
	"fmt"
	"testing"
)

/*
Prints numbers from 0 to 4 using a for range over channel loop
*/
func TestRangeOverChannel(t *testing.T) {
	resource := 0

	incoming := make(chan int)

	go func(outgoing chan<- int) {
		for {
			if resource == 5 {
				close(outgoing) // this will cause the loop below to end
				break
			}
			outgoing <- resource
			resource++
		}
	}(incoming)

	for num := range incoming {
		fmt.Println(num)
	}
}
