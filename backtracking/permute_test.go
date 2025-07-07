package backtracking

import (
	"slices"
)

func PermuteUnique(nums []int) [][]int {
	slices.Sort(nums)
	result := [][]int{}
	used := make([]bool, len(nums))

	dfsB(nums, used, []int{}, &result)

	return result
}

func dfsB(nums []int, used []bool, perm []int, result *[][]int) {
	if len(perm) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, perm)
		*result = append(*result, tmp)
		return
	}

	for i := range nums {
		if used[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		used[i] = true
		perm = append(perm, nums[i])
		dfsB(nums, used, perm, result)
		perm = perm[:len(perm)-1]
		used[i] = false
	}
}
