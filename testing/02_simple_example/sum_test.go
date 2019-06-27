package math

import (
	"testing"
)


func TestSum(t *testing.T) {
	sum := Sum([]int{2, -2, 3})
	if sum != 5 {
		t.Errorf("Wanted 5 but received %d", sum)
	}
}