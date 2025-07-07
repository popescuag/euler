package string

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	res := 0
	// last index of all characters is initialized as -1
	lastIndex := make([]int, 1000)
	for i := range len(lastIndex) {
		lastIndex[i] = -1
	}

	// Initialize start of current window
	start := 0
	fmt.Println(s)

	// Move end of current window
	for end := range n {
		// Find the last index of s[end]
		// Update starting index of current window as
		// maximum of current value of end and last index + 1
		start = int(math.Max(float64(start), float64(lastIndex[s[end]-'a']+1)))

		// Update result if we get a larger window
		res = int(math.Max(float64(res), float64(end-start+1)))

		// Update last index of s[end]
		lastIndex[s[end]-'a'] = end
		fmt.Printf("%d %d %d\n", start, end, res)
	}
	fmt.Println()
	return res
}

func LengthOfLongestSubstringN3(s string) int {
	ls := len(s)
	if ls == 0 {
		return 0
	}

	max := 0
	i := 0
	j := 1
	tmpStr := string(s[0])

	for i < ls {
		for {
			if j == ls {
				if validateString(tmpStr) && max < len(tmpStr) {
					max = len(tmpStr)
				}
				return max
			}
			tmpStr += string(s[j])
			if !validateString(tmpStr) {
				if max < len(tmpStr)-1 {
					max = len(tmpStr) - 1
					fmt.Printf("%v %d\n", tmpStr, max)
				}
				j++
				break
			}
			j++
		}
		tmpStr = tmpStr[1:]
		i = j + 1
		fmt.Printf("%v %d %d %d\n", tmpStr, max, i, j)
	}

	return max
}

func validateString(s string) bool {
	m := make(map[byte]int)
	len := len(s)
	for i := range len {
		if _, ok := m[s[i]]; ok {
			return false
		}
		m[s[i]] = 0
	}
	return true
}

func TestString(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstring(" "))
	assert.Equal(t, 1, lengthOfLongestSubstring("a"))
	assert.Equal(t, 0, lengthOfLongestSubstring(""))
	t.Fail()
}

func TestStr(t *testing.T) {
	assert.Equal(t, "bca", "abca"[1:])
}

func TestPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	assert.Equal(t, "fl", longestCommonPrefix(strs))
	strs = []string{"a"}
	assert.Equal(t, "a", longestCommonPrefix(strs))

}

func longestCommonPrefix(strs []string) string {
	prefix := ""
	if len(strs[0]) == 0 {
		return prefix
	}
	if len(strs) == 1 {
		return strs[0]
	}
	i := 0
	for {
		c := strs[0][i]
		for _, s := range strs {
			l := len(s)
			if l == 0 {
				return ""
			}
			if i == l || s[i] != c {
				return prefix
			}
		}
		prefix += string(c)
		//fmt.Println(prefix)
		i++
		if i == len(strs[0]) {
			return prefix
		}
	}
}
