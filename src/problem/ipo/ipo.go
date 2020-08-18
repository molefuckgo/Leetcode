package main

func main() {
	//k := 10
	//W := 0
	//Profits := []int{1, 2, 3}
	//Capital := []int{0, 1, 2}
	//fmt.Println(findMaximizedCapital(k, W, Profits, Capital))

}

//k=2, W=0, Profits=[1,2,3], Capital=[0,1,1].
func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	for i := 0; i < k; i++ {
		canDo := make([]int, 0)
		for j := 0; j < len(Capital); j++ {
			if Capital[j]-W <= 0 {
				canDo = append(canDo, j)
			}
		}
		if len(canDo) == 0 {
			return W
		}
		var curChoose, curChooseIdx int
		for j := 0; j < len(canDo); j++ {
			if Profits[canDo[j]] > curChoose {
				curChoose = Profits[canDo[j]]
				curChooseIdx = canDo[j]
			}
		}
		W += curChoose
		if len(Profits) > 0 {
			Profits = append(Profits[:curChooseIdx], Profits[curChooseIdx+1:]...)
			Capital = append(Capital[:curChooseIdx], Capital[curChooseIdx+1:]...)
		}

	}
	return W
}
func max2(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
