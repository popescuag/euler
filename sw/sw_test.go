package sw

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func sumK(arr []int, k int) int {
	len := len(arr)
	maxSum := 0
	sum := 0
	if len <= k {
		for i := range len {
			sum += arr[i]
		}
		return sum
	}

	for i := range k {
		sum += arr[i]
	}
	maxSum = sum

	for i := k; i < len; i++ {
		sum -= arr[i-k]
		sum += arr[i]
		fmt.Println(sum)
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func TestSumK(t *testing.T) {
	arr := []int{1, 3, 5, 2, -10}
	sum := 10
	k := 3

	assert.Equal(t, sum, sumK(arr, k))

	arr = []int{1, 3, -10}
	sum = -6
	k = 3

	assert.Equal(t, sum, sumK(arr, k))

	arr = []int{1, 3}
	sum = 4
	k = 3

	assert.Equal(t, sum, sumK(arr, k))

	arr = []int{}
	sum = 0
	k = 3

	assert.Equal(t, sum, sumK(arr, k))

}
