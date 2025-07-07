package dynamic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	result := [][]int{{1}, {1, 1}}
	for i := 3; i <= numRows; i++ {
		arr := make([]int, i)
		arr[0] = 1
		arr[i-1] = 1
		for j := 1; j < i-1; j++ {
			arr[j] = result[i-2][j-1] + result[i-2][j]
		}
		result = append(result, arr)
	}

	return result
}
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	result := make([]int, rowIndex+1)
	result[0] = 1
	result[1] = 1
	result[rowIndex] = 1
	for i := 2; i <= rowIndex+1; i++ {
		arr := make([]int, i)
		copy(arr, result[:i])
		//fmt.Println(arr)
		arr[0] = 1
		arr[i-1] = 1
		for j := 1; j < i-1; j++ {
			arr[j] += result[j-1]
		}
		//fmt.Println(arr)
		copy(result, arr)
		fmt.Println()
	}

	return result
}

func TestPascal(t *testing.T) {
	assert.Equal(t, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}, generate(5))
}

func TestGetRow(t *testing.T) {
	assert.Equal(t, []int{1, 4, 6, 4, 1}, getRow(4))
}
