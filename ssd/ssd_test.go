package ssd

import "testing"

func dif(n int) int {
	// The sum of the squares of the first n natural numbers
	// is given by the formula: n(n + 1)(2n + 1) / 6
	// The square of the sum of the first n natural numbers
	// is given by the formula: (n(n + 1) / 2) ^ 2
	sumOfSquares := n * (n + 1) * (2*n + 1) / 6
	squareOfSum := (n * (n + 1) / 2)
	squareOfSum *= squareOfSum

	return squareOfSum - sumOfSquares
}

func Test(t *testing.T) {
	expected := 25164150
	result := dif(100)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
