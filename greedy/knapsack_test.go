package greedy

import (
	"testing"
)

func TestKnapsack(t *testing.T) {
	tests := []struct {
		items     []KnapsackItem
		capacity  int
		maxProfit int
	}{
		{[]KnapsackItem{}, 0, 0},
		{[]KnapsackItem{{Value: 1, Weight: 1}}, 1, 1},
		{[]KnapsackItem{{Value: 1, Weight: 1}, {Value: 10, Weight: 5}}, 5, 10},
		{[]KnapsackItem{{Value: 1, Weight: 1}, {Value: 10, Weight: 5}}, 2, 4},
		{[]KnapsackItem{{Value: 6, Weight: 2}, {Value: 10, Weight: 2}, {Value: 12, Weight: 3}}, 5, 22},
		{[]KnapsackItem{{Value: 7, Weight: 3}, {Value: 5, Weight: 2}, {Value: 2, Weight: 1}, {Value: 10, Weight: 4}}, 6, 14},
		{[]KnapsackItem{{Value: 50, Weight: 10}, {Value: 55, Weight: 20}, {Value: 110, Weight: 30}}, 50, 187},
	}

	for i, test := range tests {
		if got := Knapsack(test.items, test.capacity); got != test.maxProfit {
			t.Fatalf("Failed test case #%d. Want %d got %v", i, test.maxProfit, got)
		}
	}
}
