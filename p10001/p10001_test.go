package p10001

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func nthPrime(n int) int {
	if n < 1 {
		return 0
	}

	count := 2
	num := 3

	for count < n {
		num += 2
		if isPrime(num) {
			count++
		}
	}

	return num
}
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func TestPrime(t *testing.T) {
	fmt.Println("Testing nthPrime function")
	expected := 104743
	result := nthPrime(10001)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	} else {
		fmt.Println("Test passed")
	}
}

func Test(t *testing.T) {
	n := 9
	r := int(math.Floor(math.Sqrt(float64(n))))
	assert.Equal(t, n, r*r)
	n = 19
	r = int(math.Floor(math.Sqrt(float64(n))))
	assert.NotEqual(t, n, r*r)
}
