package besttimetobuyandsellstock

import (
	"math/rand"
	"testing"
)

// checkStatement fails the test if the input doesn't follow the constraints
// of the problem (leetcode121.md):
//   - 1 <= prices.length <= 10^5
//   - 0 <= prices[i] <= 10^4
func checkStatement(t *testing.T, prices []int) {
	t.Helper()

	if len(prices) < 1 || len(prices) > 100_000 {
		t.Fatalf("test setup error: prices.length is %d, want between 1 and 100000", len(prices))
	}
	for _, p := range prices {
		if p < 0 || p > 10_000 {
			t.Fatalf("test setup error: %d is out of the range [0, 10000]", p)
		}
	}
}

// bruteForce is an independent way of finding the answer: it tries every
// pair of days (buy first, sell later). It is too slow for big inputs, so it
// is only used to check the test cases and small random inputs.
func bruteForce(prices []int) int {
	best := 0
	for buy := 0; buy < len(prices); buy++ {
		for sell := buy + 1; sell < len(prices); sell++ {
			best = max(best, prices[sell]-prices[buy])
		}
	}

	return best
}

// withLowestPrice calculates the same answer in O(n), remembering the lowest
// price seen so far. It is used to get the expected result of big inputs.
func withLowestPrice(prices []int) int {
	lowest := prices[0]
	best := 0
	for _, p := range prices[1:] {
		best = max(best, p-lowest)
		lowest = min(lowest, p)
	}

	return best
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"leetcode example 1", []int{7, 1, 5, 3, 6, 4}, 5},
		{"leetcode example 2", []int{7, 6, 4, 3, 1}, 0},

		// Too few days to make a transaction.
		{"a single day", []int{5}, 0},
		{"a single day with price zero", []int{0}, 0},

		// Two days.
		{"two days, the price goes up", []int{1, 2}, 1},
		{"two days, the price goes down", []int{2, 1}, 0},
		{"two days, the price doesn't change", []int{3, 3}, 0},

		// No profit.
		{"all prices are equal", []int{4, 4, 4, 4}, 0},
		{"strictly descending prices", []int{5, 4, 3, 2, 1}, 0},
		{"the lowest price is on the last day", []int{5, 4, 3, 1}, 0},
		{"descending prices with a repeated value", []int{5, 5, 3, 3, 1}, 0},

		// The profit is the last price minus the first one.
		{"strictly ascending prices", []int{1, 2, 3, 4, 5}, 4},
		{"ascending prices with a repeated value", []int{1, 1, 2, 2, 3}, 2},

		// The lowest and highest prices are not the best pair.
		{"the highest price comes before the lowest one", []int{2, 10, 1, 5}, 8},
		{"the lowest price is not the best day to buy", []int{3, 2, 6, 5, 0, 3}, 4},
		{"the highest price is on the first day", []int{10, 5, 3, 7}, 4},
		{"the best profit is at the start", []int{1, 9, 2, 3, 4, 5}, 8},
		{"the best profit is at the end", []int{1, 3, 2, 8, 4, 9}, 8},
		{"a deeper dip appears later", []int{5, 1, 3, 0, 10}, 10},
		{"the price drops after the best sale", []int{2, 4, 1}, 2},

		// Zero prices.
		{"free stock", []int{0, 5, 0, 3}, 5},
		{"free stock and then the maximum price", []int{0, 10000}, 10000},

		// Values at the constraint bounds.
		{"maximum price and then zero", []int{10000, 0}, 0},
		{"zero, maximum price and zero again", []int{0, 10000, 0}, 10000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.prices)

			// Check that the expected result of the test case is right.
			if want := bruteForce(tt.prices); want != tt.want {
				t.Fatalf("test setup error: the expected result is %d, but the right answer for %v is %d", tt.want, tt.prices, want)
			}

			if got := MaxProfit(tt.prices); got != tt.want {
				t.Errorf("MaxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}

// TestMaxProfitRandomSmallInputs compares the result with a brute force
// calculation for many small random inputs. It uses a fixed seed so a
// failure can be reproduced.
func TestMaxProfitRandomSmallInputs(t *testing.T) {
	r := rand.New(rand.NewSource(1))

	for range 1000 {
		n := 1 + r.Intn(15) // 1 to 15 days
		prices := make([]int, n)
		for i := range prices {
			prices[i] = r.Intn(21) // 0 to 20
		}
		checkStatement(t, prices)

		want := bruteForce(prices)
		if got := MaxProfit(prices); got != want {
			t.Fatalf("MaxProfit(%v) = %d, want %d", prices, got, want)
		}
	}
}

// TestMaxProfitLargeInput uses the maximum allowed length (10^5). A solution
// that tries every pair of days (O(n^2)) takes very long here.
func TestMaxProfitLargeInput(t *testing.T) {
	const n = 100_000
	r := rand.New(rand.NewSource(2))

	allEqual := make([]int, n)
	descending := make([]int, n)
	ascending := make([]int, n)
	alternating := make([]int, n)
	random := make([]int, n)
	for i := range n {
		allEqual[i] = 5_000
		descending[i] = 10_000 - i/10 // 10000 down to 1
		ascending[i] = i / 10         // 0 up to 9999
		if i%2 == 1 {
			alternating[i] = 10_000
		}
		random[i] = r.Intn(10_001)
	}

	tests := []struct {
		name   string
		prices []int
		want   int // -1 means: use the O(n) calculation
	}{
		{"all prices are equal", allEqual, 0},
		{"descending prices", descending, 0},
		{"ascending prices", ascending, 9_999},
		{"prices alternating between zero and the maximum", alternating, 10_000},
		{"random prices", random, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkStatement(t, tt.prices)

			want := withLowestPrice(tt.prices)
			if tt.want != -1 && tt.want != want {
				t.Fatalf("test setup error: the expected result is %d, but the right answer is %d", tt.want, want)
			}

			if got := MaxProfit(tt.prices); got != want {
				t.Errorf("MaxProfit(large input) = %d, want %d", got, want)
			}
		})
	}
}
