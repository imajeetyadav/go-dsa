package greedy

import (
	"testing"
)

func TestJumpGame(t *testing.T) {
	tests := []struct {
		jumps          []int
		canReachTheEnd bool
	}{
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{1, 0, 1}, false},
		{[]int{3, 1, 0, 0, 1}, false},
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{2, 3, 1, 1, 2, 0, 1}, true},
	}

	for i, test := range tests {
		if got := JumpGame(test.jumps); got != test.canReachTheEnd {
			t.Fatalf("Failed test case #%d. Want %t got %t", i, test.canReachTheEnd, got)
		}
	}
}
