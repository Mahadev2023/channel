package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send value
	go send(eve, odd, quit)

	//receive
	receive(eve, odd, quit)

	fmt.Println("about to close")

}
func receive(e, o <-chan int, q <-chan int) {
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

func send(e, o chan<- int, q chan<- int) {
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
