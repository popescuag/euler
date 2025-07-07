package backtracking

import (
	"strings"
)

func LetterCasePermutation(s string) []string {
	result := []string{}

	var dfs func(index int, str string)

	dfs = func(index int, str string) {
		if index == len(str) {
			result = append(result, str)
			return
		}

		if str[index] >= '0' && str[index] <= '9' {
			dfs(index+1, str)
			return
		}

		newStr := str[:index]
		newStr += strings.ToUpper(string(str[index]))
		newStr += str[index+1:]

		dfs(index+1, newStr)

		newStr = str[:index]
		newStr += strings.ToLower(string(str[index]))
		newStr += str[index+1:]

		dfs(index+1, newStr)
	}

	dfs(0, s)

	return result
}
