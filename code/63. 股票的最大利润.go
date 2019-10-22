package code

func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0] // 记录之前的最小价值
	maxProfit := 0
	for _, price := range prices {
		maxProfit = MaxInt(maxProfit, price-minPrice)
		minPrice = MinInt(minPrice, price)
	}
	return maxProfit
}
