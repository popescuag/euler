package backtracking

import (
	"slices"
)

func Subsets(nums []int) [][]int {
	result := [][]int{{}}
	l := len(nums)
	for _, n := range nums {
		result = append(result, []int{n})
	}
	if l == 1 {
		return result
	}
	slices.Sort(nums)
	for j := 2; j <= l; j++ {
		dfs3(nums, j, 0, []int{}, &result)
	}
	return result
}

func dfs3(nums []int, k int, index int, combo []int, result *[][]int) {
	if len(combo) == k {
		tmp := make([]int, len(combo))
		copy(tmp, combo)
		*result = append(*result, tmp)
		return
	}

	for i := index; i < len(nums); i++ {
		if len(combo) > 0 && nums[i] <= combo[len(combo)-1] {
			continue
		}
		combo = append(combo, nums[i])
		dfs3(nums, k, index+1, combo, result)
		combo = combo[:len(combo)-1]
	}
}
