package dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func fib(n int) int {
	n0 := 0
	n1 := 1

	switch n {
	case 0:
		return n0
	case 1:
		return n1
	}

	result := 0
	idx := 2

	for {
		result = n0 + n1
		if idx == n {
			break
		}
		idx++
		n0, n1 = n1, result
	}

	return result
}

func tribonacci(n int) int {
	n0 := 0
	n1 := 1
	n2 := 1

	switch n {
	case 0:
		return n0
	case 1:
		return n1
	case 2:
		return n2
	}

	result := 0
	idx := 3

	for {
		result = n0 + n1 + n2
		if idx == n {
			break
		}
		idx++
		n0, n1, n2 = n1, n2, result
	}

	return result
}

func TestFibo(t *testing.T) {
	assert.Equal(t, 3, fib(4))
}

func TestTribo(t *testing.T) {
	assert.Equal(t, 1389537, tribonacci(25))
}
