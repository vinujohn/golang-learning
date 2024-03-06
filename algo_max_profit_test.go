package learning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// sliding window problem
// keep track of the min so as we go forward
func maxProfit(prices []int) int {
	profit, min := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if (prices[i] - min) > profit {
			profit = (prices[i] - min)
		}
	}
	return profit
}

func TestMaxProfit(t *testing.T) {
	assert.Equal(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
	assert.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))
}
