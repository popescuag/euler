package trbn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Trbn(n int) int {
	t0 := 0
	t1 := 1
	t2 := 1
	c := 3
	t3 := 0

	for c <= n {
		t3 = t0 + t1 + t2
		t0, t1, t2 = t1, t2, t3
		fmt.Printf("t0: %d, t1: %d, t2: %d, t3: %d\n", t0, t1, t2, t3)
		c++
	}
	return t3
}

func Test(t *testing.T) {
	assert.Equal(t, 3136, Trbn(15), "")
	assert.Equal(t, 2555757, Trbn(26), "")
}
