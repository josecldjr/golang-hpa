package main

import (
	"math"
)

/*
SquareRoot - aplly a square root function over a number, using a loop and return the result
*/
func SquareRoot() float64 {

	x := 0.0001

	for i := 0; i < 1000000; i++ {
		x += math.Sqrt(x)
	}

	return x
}
