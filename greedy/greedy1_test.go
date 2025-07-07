package greedy

func GroupThePeople(groupSizes []int) [][]int {
	groupings := make(map[int][][]int)
	result := [][]int{}
	for i, s := range groupSizes {
		groups := groupings[s]
		l := len(groups)
		if l == 0 || len(groups[l-1]) == s {
			groups = append(groups, []int{})
			l++
		}
		groups[l-1] = append(groups[l-1], i)

		groupings[s] = groups
	}

	for _, v := range groupings {
		result = append(result, v...)
	}

	return result
}
