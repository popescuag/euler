package dynamic

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func maximalSquare(matrix [][]byte) int {
	rc, cc := len(matrix), len(matrix[0])
	cache := make(map[string]int)

	var dfs func(i, j int) int
	max := -1

	dfs = func(i, j int) int {
		if i >= rc || j >= cc {
			return 0
		}
		key := strconv.Itoa(i) + strconv.Itoa(j)
		if _, ok := cache[key]; !ok {
			d := float64(dfs(i+1, j))
			r := float64(dfs(i, j+1))
			diag := float64(dfs(i+i, j+1))

			cache[key] = 0
			if matrix[i][j] == '1' {
				cache[key] = 1 + int(math.Min(d, math.Min(r, diag)))
			}
		}

		if cache[key] > max {
			max = cache[key]
		}

		return cache[key]
	}
	dfs(0, 0)

	fmt.Println(cache)

	return max * max
}

func TestMax(t *testing.T) {
	matrix := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '1', '1', '0'}, {'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'}, {'0', '0', '1', '1', '1'}}

	assert.Equal(t, 16, maximalSquare(matrix))
}
