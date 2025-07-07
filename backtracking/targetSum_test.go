package backtracking

import (
	"sort"
	"strings"
	"testing"

	"fmt"

	"slices"
	"strconv"

	"github.com/stretchr/testify/assert"
)

func findTargetSumWays(nums []int, target int) int {
	s := ""
	return findTarget(nums, target, 0, &[]string{}, &s)
}

func findTarget(nums []int, target int, index int, solutions *[]string, currentSolution *string) int {
	//fmt.Printf("%v %v\n", nums, index)
	if index == len(nums) {
		if verifySolution(nums, target, solutions, *currentSolution) {
			*solutions = append(*solutions, *currentSolution)
			fmt.Printf("Solution found: %v %v\n", *currentSolution, nums)
			*currentSolution = ""
			return 1
		}
		*currentSolution = ""
		return 0
	}
	count := 0
	if index == 0 {
		*currentSolution = strings.Repeat("+", len(nums))
		if verifySolution(nums, target, solutions, *currentSolution) {
			*solutions = append(*solutions, *currentSolution)
			fmt.Printf("Solution (0) found: %v\n", *currentSolution)
			count++
		}
		*currentSolution = ""
	}
	for i := index; i < len(nums); i++ {
		nums[i] *= (-1)
		*currentSolution += strconv.Itoa(nums[i])
		fmt.Printf("Before: %v %v %d\n", nums, *currentSolution, index)
		count += findTarget(nums, target, index+1, solutions, currentSolution)
		//fmt.Printf("After: %v\n", *currentSolution)
		if *currentSolution != "" {
			*currentSolution = (*currentSolution)[:len((*currentSolution))-1]
		}
		*currentSolution += "+"
		nums[i] *= (-1)
	}
	return count
}

func verifySolution(nums []int, target int, solutions *[]string, currentSolution string) bool {
	s := 0
	for _, e := range nums {
		s += e
	}
	if s != target {
		return false
	}
	//Check if the solution was not previously found
	return !slices.Contains(*solutions, currentSolution)
}

func TestTargetSum(t *testing.T) {
	input := [][]int{{1, 1, 1, 1, 1}, {1}, {1}, {1, 0}}
	target := []int{3, 1, 2, 1}
	expected := []int{5, 1, 0, 2}

	assert.Equal(t, expected[0], findTargetSumWays(input[0], target[0]))

	// for i := range input {
	// 	assert.Equal(t, expected[i], findTargetSumWays(input[i], target[i]))
	// }
	//t.Fail()
}

func TestSlice(t *testing.T) {
	s := "abcde"
	s = s[:len(s)-1]
	assert.Equal(t, "abcd", s)
}

func TestKnapsack(t *testing.T) {
	assert.Equal(t, 5, knapsack([]int{1, 1, 1, 1, 1}, 4))
	t.Fail()
}

func TestSort(t *testing.T) {
	arr := [][]int{{2, 3}, {4, 1}, {5, 7}}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})

	fmt.Println(arr)
	t.Fail()
}

func knapsack(nums []int, target int) int {
	// dp[i] := the number of ways to sum to i by nums so far
	dp := make([]int, target+1)
	dp[0] = 1

	for _, num := range nums {
		fmt.Println(dp)
		for i := target; i >= num; i-- {
			dp[i] += dp[i-num]
		}
		fmt.Println(dp)
		fmt.Println()
	}
	return dp[target]
}
