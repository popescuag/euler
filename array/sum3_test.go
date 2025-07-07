package array

import "slices"

func Sum3(numbers []int, sum int) [][]int {
	result := [][]int{}
	slices.Sort(numbers)

	for i, e := range numbers {
		if i > 0 && e == numbers[i-1] {
			continue
		}
		l, r := i+1, len(numbers)-1
		for l < r {
			s := e + numbers[l] + numbers[r]
			if s > sum {
				r--
			} else if s < sum {
				l++
			}
			if s == sum {
				result = append(result, []int{e, numbers[l], numbers[r]})
				l++
				for numbers[l-1] == numbers[l] && l < r {
					l++
				}
			}
		}
	}

	return result
}
