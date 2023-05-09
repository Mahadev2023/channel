package main

import "fmt"

func main() {

	// c := make(chan  int, 2)

	//send value channel
	c := make(chan<- int, 2)

	c <- 32
	c <- 43

	//by send only channel we cannot receive form it
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	fmt.Println("###########")

	fmt.Printf("%T\n", c)
}
