package array

func TwoSumIIBF(numbers []int, target int) []int {
	len := len(numbers)
	for i := range len {
		for j := i + 1; j < len; j++ {
			sum := numbers[i] + numbers[j]
			if sum == target {
				return []int{i + 1, j + 1}
			} else if sum > target {
				break
			}
		}
	}
	return []int{}
}

func TwoSumII(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l < r {
		currSum := numbers[l] + numbers[r]
		if currSum == target {
			return []int{l + 1, r + 1}
		}
		if currSum > target {
			r--
		} else {
			l++
		}
	}
	return []int{}
}
