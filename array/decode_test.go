package array

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func decode(s string) string {
	retStr := ""

	for i := 0; i < len(s); {
		nStr := ""
		if s[i] >= '0' && s[i] <= '9' {
			for s[i] >= '0' && s[i] <= '9' {
				nStr += string(s[i])
				i++
			}
			n, _ := strconv.Atoi(nStr)
			j := i + 1
			op := 0
			for ; j < len(s); j++ {
				//fmt.Printf("%c\n", s[j])
				if s[j] == ']' {
					if op == 0 {
						break
					}
					op--
				}
				if s[j] == '[' {
					op++
				}
			}
			//fmt.Printf("Repeat %v %v times\n", s[(i+1):j], n)
			retStr += strings.Repeat(decode(s[(i+1):j]), n-1)
		} else if s[i] == '[' || s[i] == ']' {
			i++
			continue
		} else {
			retStr += string(s[i])
			i++
		}
	}
	return retStr
}

func TestDecode(t *testing.T) {
	input := []string{"10[a]", "abc", "[]", "3[a]b", "b[]a", "b2[a]", "2[3[a]bc]", "[]a"}
	output := []string{"aaaaaaaaaa", "abc", "", "aaab", "ba", "baa", "aaabcaaabc", "a"}

	//assert.Equal(t, output[3], decode(input[3]))

	for i := range input {
		assert.Equal(t, output[i], decode(input[i]))
	}
}
