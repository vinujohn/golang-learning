package learning

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func minCoins(denominations []int, amount int) int {
	// Create a slice to store the minimum number of coins needed for each amount
	minCoins := make([]int, amount+1)

	// Initialize the slice with a large number (in this case, math.MaxInt32)
	for i := range minCoins {
		minCoins[i] = math.MaxInt32
	}

	// We can make 0 amount with 0 coins
	minCoins[0] = 0

	// For each denomination, calculate the minimum number of coins needed for each amount
	for _, denomination := range denominations {
		for i := denomination; i <= amount; i++ {
			// Check if we can make the amount with this denomination
			if minCoins[i-denomination]+1 < minCoins[i] {
				minCoins[i] = minCoins[i-denomination] + 1
			}
		}
	}

	// If we can't make the amount, return -1
	if minCoins[amount] == math.MaxInt32 {
		return -1
	}

	// Return the minimum number of coins needed to make the amount
	return minCoins[amount]
}

func TestMinCoins(t *testing.T) {
	tests := []struct {
		name          string
		denominations []int
		amount        int
		expected      int
	}{
		{
			name:          "Single denomination",
			denominations: []int{1},
			amount:        10,
			expected:      10,
		},
		{
			name:          "Multiple denominations",
			denominations: []int{1, 2, 5},
			amount:        11,
			expected:      3,
		},
		{
			name:          "No solution",
			denominations: []int{2},
			amount:        3,
			expected:      -1,
		},
		{
			name:          "Zero amount",
			denominations: []int{1, 2, 5},
			amount:        0,
			expected:      0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := minCoins(tt.denominations, tt.amount)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
