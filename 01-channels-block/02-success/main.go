package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	// c <- 42

	gor1 := func() {
		fmt.Println("Hello From gor1")
		time.Sleep(time.Second * 3)
		c <- 42
		fmt.Println("Hello form gor1 after c wrote")
		// time.Sleep(time.Second * 3)
	}
	go gor1()
	fmt.Println("HEllo before c read")

	fmt.Println(<-c)

	fmt.Println("HEllo after c read")
}
