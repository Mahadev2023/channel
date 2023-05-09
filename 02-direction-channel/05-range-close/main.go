package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	//send
	// go foo(c)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) //if we do not close c then deadlock is happen(because below receiver is waiting)

	}()

	//receive//
	// go bar(c)
	// bar(c)

	for v := range c {
		fmt.Println(v)
	}

	// time.Sleep(time.Second * 2)
	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)

}

// receive
// func bar(c <-chan int) {
// 	fmt.Println("value in bar", <-c)
// }
