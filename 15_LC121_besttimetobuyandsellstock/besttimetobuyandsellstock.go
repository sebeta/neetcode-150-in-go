// Package besttimetobuyandsellstock solves the "Best time to buy and sell stock" problem.
package besttimetobuyandsellstock

// MaxProfit returns the maximum profit that can be made by buying a stock on
// one day and selling it on a later day, where prices[i] is the price of the
// stock on day i. It returns 0 if no profit is possible. prices must have at
// least one element.
//
// It goes through the prices once, keeping the lowest price seen so far
// (minPrice). Each day is considered as the day to sell: the profit is its
// price minus minPrice, which comes from earlier days, and the best profit
// found is kept in maxProfit. The profit is calculated before updating
// minPrice, so the stock is never bought and sold on the same day.
//
// Time complexity: O(n), where n is the length of prices.
// Space complexity: O(1), since only two variables are used.
//
// NOTE: This was my solution. It is the same as in the video.
func MaxProfit(prices []int) int {
	var maxProfit int
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - minPrice
		maxProfit = max(maxProfit, profit)
		minPrice = min(minPrice, prices[i])
	}

	return maxProfit
}
