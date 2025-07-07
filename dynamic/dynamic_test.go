package dynamic

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func lis(arr []int) int {
	fmt.Println(arr)
	mem := make([]int, len(arr))

	for i := range len(arr) {
		mem[i] = 1
	}

	i := 0
	j := 1

	for ; j < len(arr); j++ {
		for i = range j {
			if arr[i] < arr[j] {
				if mem[i]+1 > mem[j] {
					mem[j] = mem[i] + 1
				}
			}
		}
	}
	maxLIS := mem[0]
	for i := 1; i < len(arr); i++ {
		if mem[i] > maxLIS {
			maxLIS = mem[i]
		}
	}
	fmt.Println(mem)

	return maxLIS
}

func CanJump(nums []int) bool {
	l := len(nums)
	i := 0
	for reach := 0; i < l && i <= reach; i++ {
		reach = int(math.Max(float64(reach), float64(i+nums[i])))
		//fmt.Printf("%v %v\n", reach, i)
	}

	return i == l
}

func TestLIS(t *testing.T) {
	arr := []int{15,
		-1,
		8,
		2,
		5,
		11,
		7,
		10}
	assert.Equal(t, 5, lis((arr)))
	assert.Equal(t, 6, lis([]int{15,
		-1,
		8,
		2,
		5,
		11,
		7,
		10,
		20}))
}
