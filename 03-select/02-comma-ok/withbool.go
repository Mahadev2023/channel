package main

import "fmt"

func main1() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	//send value
	go send(eve, odd, quit)

	//receive
	receive(eve, odd, quit)

	fmt.Println("about to close")

}
func receive1(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the eve channel:", v)
		case v := <-o:
			fmt.Println("From the odd channel:", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("From comma ok", i, ok)
				return
			} else {
				fmt.Println("From comma ok", i)
			}

		}

	}
}

func send1(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	// close(e)
	// close(o)
	close(q)
}
