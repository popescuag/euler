package backtracking

func Combine(n int, k int) [][]int {
	if n == 1 && k == 1 {
		return [][]int{{1}}
	}

	c := []int{}
	result := [][]int{}

	dfsC(n, k, 1, c, &result)

	return result
}

func dfsC(n int, k int, index int, combo []int, result *[][]int) {
	if len(combo) == k {
		//fmt.Println(combo)
		tmp := make([]int, len(combo))
		copy(tmp, combo)
		*result = append(*result, tmp)
		return
	}

	for i := index; i <= n; i++ {
		if len(combo) > 0 && i <= combo[len(combo)-1] {
			continue
		}
		combo = append(combo, i)
		dfsC(n, k, index+1, combo, result)
		combo = combo[:len(combo)-1]
	}
}
