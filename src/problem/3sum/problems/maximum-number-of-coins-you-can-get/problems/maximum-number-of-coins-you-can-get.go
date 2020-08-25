package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxCoins([]int{}))
}

func maxCoins(piles []int) int {
	sort.Ints(piles)
	pilesLen := len(piles)
	times := pilesLen / 3
	result := 0
	for i := pilesLen - 2; ; i -= 2 {
		if times == 0 {
			break
		}
		result += piles[i]
		times--
	}
	return result
}
