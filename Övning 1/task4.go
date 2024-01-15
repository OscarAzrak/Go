package main

import (
	"fmt"
)

//Add sums the values in a
func Add(a []int, res chan<- int) {
	// TODO

	sum := 0
	for i := range a {
		sum += a[i]
	}
	res <- sum

}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	n := len(a)
	ch := make(chan int)
	go Add(a[:n/2], ch)
	go Add(a[n/2:], ch)

	// TODO: Get the subtotals from the channel and print their sum.
	sum1 := <-ch
	sum2 := <-ch

	close(ch)

	totSum := sum1 + sum2

	fmt.Println("Sum of half of list: ", sum1)
	fmt.Println("Sum of second half of list: ", sum2)
	fmt.Println("Total sum: ", totSum)
}
