func maxProfit(prices []int) int {
	maxProfit, buy := 0, 0
	for sell := 1; sell < len(prices); sell++ {
		if prices[buy] > prices[sell] {
			buy = sell
			continue
		}

		profit := prices[sell] - prices[buy]
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
