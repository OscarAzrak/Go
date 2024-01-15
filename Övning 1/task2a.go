package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	var temp float64 //create temporary variable
	for {
		temp = z
		z = z - (z*z-x)/(2*z)

		if math.Abs(temp-z) < 1e-6 { //If error is less than 1/1000000, loop will stop
			break
		}
	}
	return z
}

func main() {
	estimation := Sqrt(2)
	real := math.Sqrt(2)
	fmt.Print("Guess: ", estimation, ", Expected: ", real, ", Error: ", math.Abs(estimation-real), "\n")

}
