package main

func main() {

}

func maxProfit(prices []int) int {
	pricesLen := len(prices)
	if pricesLen == 0 {
		return 0
	}
	minBuy := prices[0]
	result := 0
	for i := 1; i < pricesLen; i++ {
		if prices[i] > minBuy {
			result = max2(prices[i]-minBuy, result)
		} else {
			minBuy = prices[i]
		}
	}
	return result
}

func max2(x, y int) int {
	if x <= y {
		return y
	}
	return x
}
