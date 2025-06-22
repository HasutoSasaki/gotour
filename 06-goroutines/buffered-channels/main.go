package main

import "fmt"

func main() {
	ch := make(chan int, 100)
	ch <- 1
	ch <- 2
	ch <- 3           // this will block if the channel is full
	fmt.Println(<-ch) // prints 1
	fmt.Println(<-ch) // prints 2
}
