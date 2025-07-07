package backtracking

import (
	"fmt"
	"testing"
)

func GeneratePermutationsRecursive(numbers []int) [][]int {
	if len(numbers) == 0 {
		return [][]int{}
	}

	if len(numbers) == 1 {
		return [][]int{{numbers[0]}}
	}

	permutations := [][]int{}

	for i, num := range numbers {
		remaining := make([]int, len(numbers)-1)
		copy(remaining[:i], numbers[:i])
		copy(remaining[i:], numbers[i+1:])

		subPermutations := GeneratePermutationsRecursive(remaining)

		for _, p := range subPermutations {
			permutations = append(permutations, append([]int{num}, p...))
		}
	}

	return permutations
}

func GetPermutations(idx int, arr []int, result [][]int) {
	if idx == len(arr) {
		fmt.Printf("%v \n", arr)
		return
	}

	for i := idx; i < len(arr); i++ {
		arr[idx], arr[i] = arr[i], arr[idx]
		GetPermutations(idx+1, arr, result)
		arr[idx], arr[i] = arr[i], arr[idx]
	}
}

func TestPerm(t *testing.T) {
	arr := []int{1, 2, 3}
	result := [][]int{}
	GetPermutations(0, arr, result)
	fmt.Println(result)
	t.Fail()
}
