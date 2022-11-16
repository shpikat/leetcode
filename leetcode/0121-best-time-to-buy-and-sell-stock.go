package leetcode

func maxProfit(prices []int) int {
	profit := 0
	bought := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < bought {
			bought = prices[i]
		} else {
			current := prices[i] - bought
			if current > profit {
				profit = current
			}
		}
	}
	return profit
}
