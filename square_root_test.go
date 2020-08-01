package main

import (
	"math"
	"testing"
)

func TestSquareRoot(t *testing.T) {

	var expectedValue float64 = 0.0001

	for i := 0; i < 1000000; i++ {
		expectedValue += math.Sqrt(expectedValue)
	}

	resultValue := SquareRoot()

	if expectedValue != resultValue {
		t.Errorf("The returned value of the function is wrong ! Expected %v , received %v", expectedValue, resultValue)
	}
}
