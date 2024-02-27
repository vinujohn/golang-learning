package learning

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
Prints numbers from 0 to 4 then 0 using a fixed iteration for loop
*/
func TestForOverChannel(t *testing.T) {
	resource := 0

	incoming := make(chan int)

	go func(outgoing chan<- int) {
		for {
			if resource == 5 {
				close(outgoing) // this will cause the loop below to print 0
				break
			}
			outgoing <- resource
			resource++
		}
	}(incoming)

	// prints 0 in the last iteration because reading from a closed channel
	// reads the default zero value of the channel
	for i := 0; i < 6; i++ {
		fmt.Println(<-incoming)
	}
}

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

/*
Prints "default" every 500 milliseconds, "one" once after a second, and "two" once after 2 seconds.
*/
func TestForSelectChannel(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

loop:
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
			break loop
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("default")
		}

	}
}

// Prints "How are you" by taking 3 channels and putting them into a common one
func TestFanInExample(t *testing.T) {
	generate := func(data string) <-chan string {
		c := make(chan string)

		go func(incoming chan string) {
			c <- data
			time.Sleep(time.Millisecond * 100)
		}(c)

		return c
	}

	c1 := generate("How")
	c2 := generate("are")
	c3 := generate("you?")

	fanin := make(chan string)

	go func() {
		for {
			select {
			case str1 := <-c1:
				fanin <- str1
			case str2 := <-c2:
				fanin <- str2
			case str3 := <-c3:
				fanin <- str3
			}
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println(<-fanin)
	}
}

// Prints 5 numbers in random order from 5 or less goroutines
func TestFanoutExample(t *testing.T) {
	// Create a buffered channel to fan out messages to multiple goroutines
	messages := make(chan string, 10)

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create 5 consumer goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for message := range messages {
				fmt.Printf("Received message from go routine %d '%s'\n", id, message)
			}
		}(i)
	}

	// Create a producer goroutine
	go func() {
		messages <- "Hello"
		messages <- "World"
		messages <- "How"
		messages <- "Are"
		messages <- "You"
		close(messages)
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}

// go test --race -run TestRaceCondition
func TestRaceCondition(t *testing.T) {
	done := make(chan bool)

	m := make(map[string]string)
	m["name"] = "world"

	go func() {
		m["name"] = "data race"
		done <- true
	}()

	fmt.Println("Hello,", m["name"])

	<-done
}
