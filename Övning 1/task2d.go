package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

func fibonacci() func() int {
	prevNumb := 0
	currNumb := 1

	fmt.Println(prevNumb)
	fmt.Println(currNumb)

	return func() int {
		//save the current value in a holder, in order to set it correctly to previous (after)
		currTemp := currNumb
		currNumb = currNumb + prevNumb
		prevNumb = currTemp
		return currNumb
	}
}

func main() {
	fib := fibonacci()
	max := 10
	for i := 0; i < max-2; i++ {
		fmt.Println(fib())
	}
}
