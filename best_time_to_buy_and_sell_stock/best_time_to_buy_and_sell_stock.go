package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	left := 0
	right := 1

	r := 0

	for right < len(prices) {
		if prices[right] > prices[left] {
			if (prices[right] - prices[left]) > r {
				r = prices[right] - prices[left]
			}
		} else {
			left = right
		}
		right++
	}

	return r
}
