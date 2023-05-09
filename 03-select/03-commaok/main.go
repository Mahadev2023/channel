package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
		close(c)
		// time.Sleep(time.Second * 1)

		fmt.Println("channel closed")
	}()

	//this first value
	v, ok := <-c

	fmt.Println(v, ok)
	time.Sleep(time.Second * 3)
	//this read second value form channel when channel is closed
	v, ok = <-c

	fmt.Println(v, ok)

	fmt.Println("Hello about to end")
}
