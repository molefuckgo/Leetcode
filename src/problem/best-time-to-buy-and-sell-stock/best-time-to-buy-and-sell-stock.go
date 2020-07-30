package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	pricesLen := len(prices)
	if pricesLen <= 1 {
		return 0
	}
	minPrice := math.MaxInt64
	maxprofit := 0
	for i := 0; i < pricesLen; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxprofit {
			maxprofit = prices[i] - minPrice
		}
	}
	return maxprofit
}
