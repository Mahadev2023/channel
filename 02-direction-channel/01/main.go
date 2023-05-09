package main

import "fmt"

func main() {

	c := make(chan int, 2) //both send & receive channel
	c2 := make(<-chan int) //receive channel
	c3 := make(chan<- int) //send channel

	c <- 32
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("###########")

	fmt.Printf("%T\n", c)  //chan int
	fmt.Printf("%T\n", c2) //<-chan int
	fmt.Printf("%T\n", c3) //chan<- int
}
