package main

import "fmt"

func main() {
	fmt.Println(minCost([][]int{
		{17, 2, 17},
		{16, 16, 5},
		{14, 3, 19},
	}))
}

func minCost(costs [][]int) int {
	rows := len(costs)
	if rows == 0 {
		return 0
	}
	for i := rows - 2; i >= 0; i-- {
		costs[i][0] += min2(costs[i+1][1], costs[i+1][2])
		costs[i][1] += min2(costs[i+1][0], costs[i+1][2])
		costs[i][2] += min2(costs[i+1][0], costs[i+1][1])
	}
	return min2((min2(costs[0][0], costs[0][1])), costs[0][2])
}

func min2(x, y int) int {
	if x >= y {
		return y
	}
	return x
}
