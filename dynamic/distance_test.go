package dynamic

import (
	"math"
)

func MinDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)
	c := [][]int{}
	for range l1 + 1 {
		arr := make([]int, l2+1)
		for j := range l2 {
			arr[j] = math.MaxInt32
		}
		c = append(c, arr)
	}

	for j := l2; j >= 0; j-- {
		c[l1][j] = l2 - j
	}

	for i := l1; i >= 0; i-- {
		c[i][l2] = l1 - i
	}

	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			if word1[i] == word2[j] {
				c[i][j] = c[i+1][j+1]
			} else {
				c[i][j] = 1 + int(math.Min(float64(c[i+1][j]), math.Min(float64(c[i][j+1]), float64(c[i+1][j+1]))))
			}
		}
	}

	return c[0][0]
}
