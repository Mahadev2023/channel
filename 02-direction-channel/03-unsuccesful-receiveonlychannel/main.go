package main

import "fmt"

func main() {

	// c := make(chan int, 2)
	c := make(<-chan int, 2)

	//invalid operation: cannot send to receive-only channel c (variable of type <-chan int)
	// c <- 32
	// c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("###########")

	fmt.Printf("%T\n", c)
}
