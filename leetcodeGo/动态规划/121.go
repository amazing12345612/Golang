package dynamicProgramming

func maxProfit(prices []int) int {

	Min := prices[0]
	maxprofit := 0
	profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < Min {
			Min = prices[i]
		} else {
			profit = prices[i] - Min
			if maxprofit < profit {
				maxprofit = profit
			}
		}
	}
	return maxprofit

}
