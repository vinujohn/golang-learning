package learning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// sliding window problem
/*
- We initialize `minPrice` to the first price in the array and `maxProfit` to 0.
- We then iterate over the prices array. For each price:
    - If the current price is less than `minPrice`, we update `minPrice` to the current price. This is because we've found a new minimum price that could potentially lead to a larger profit.
    - If the difference between the current price and `minPrice` is greater than `maxProfit`, we update `maxProfit` to this difference. This is because we've found a new maximum profit.
- At the end of the iteration, `maxProfit` will contain the maximum profit we could have made by buying at a certain price and then selling at a later price.
*/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}

func TestMaxProfit(t *testing.T) {
	assert.Equal(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
	assert.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))
}
