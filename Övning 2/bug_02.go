package main

import "fmt"

func main() {
	ch := make(chan int)
	quit := Print(ch)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
	<-quit
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) chan struct{} {
	quit := make(chan struct{})
	go func() {
		for n := range ch { // reads from channel until it's closed
			fmt.Println(n)
		}
		close(quit)
	}()
	return quit
}