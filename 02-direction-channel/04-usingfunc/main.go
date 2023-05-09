package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	//send
	go foo(c)

	//receive
	// go bar(c)
	bar(c)

	// time.Sleep(time.Second * 2)
	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println("value in bar", <-c)
}
