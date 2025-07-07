package m35

import (
	"fmt"
	"testing"
)

func sum(limit int) int {
	sum := 0
	i := 3
	j := 5
	for i < limit {
		sum += i
		i += 3
	}
	fmt.Println("i:", i, "j:", j, "sum:", sum)
	for j < limit {
		if j%3 != 0 {
			sum += j
		}
		j += 5
	}
	fmt.Println("i:", i, "j:", j, "sum:", sum)

	return sum
}

func TestM35(t *testing.T) {
	expected := 233168
	result := sum(1000)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
