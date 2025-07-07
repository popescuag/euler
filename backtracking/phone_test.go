package backtracking

func LetterCombinations(digits string) []string {
	letters := make(map[byte][]string)
	letters[50] = []string{"a", "b", "c"}
	letters[51] = []string{"d", "e", "f"}
	letters[52] = []string{"g", "h", "i"}
	letters[53] = []string{"j", "k", "l"}
	letters[54] = []string{"m", "n", "o"}
	letters[55] = []string{"p", "q", "r", "s"}
	letters[56] = []string{"t", "u", "v"}
	letters[57] = []string{"w", "x", "y", "z"}

	l := len(digits)
	if l == 0 {
		return []string{}
	}
	if l == 1 {
		return letters[digits[0]]
	}
	results := []string{}
	dfs(digits, 0, "", letters, &results)
	return results
}

func dfs(digits string, index int, s string, lettersMap map[byte][]string, results *[]string) {
	if len(s) == len(digits) {
		*results = append(*results, s)
		return
	}
	for _, l := range lettersMap[digits[index]] {
		s += string(l)
		dfs(digits, index+1, s, lettersMap, results)
		s = s[:len(s)-1]
	}
}
