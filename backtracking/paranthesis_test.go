package backtracking

import (
	"fmt"
	"testing"
)

func GenerateParenthesis(n int) []string {
	response := []string{}
	dfs1(n, n, "", &response)

	return response
}

func dfs1(l int, r int, s string, response *[]string) {
	if l == 0 && r == 0 {
		*response = append(*response, s)
		return
	}
	if l > 0 {
		s += "("
		dfs1(l-1, r, s, response)
		s = s[:len(s)-1]
	}
	if l < r {
		s += ")"
		dfs1(l, r-1, s, response)
	}
}

func TestParanthesis(t *testing.T) {
	fmt.Println(GenerateParenthesis(4))

	t.Fail()

}
