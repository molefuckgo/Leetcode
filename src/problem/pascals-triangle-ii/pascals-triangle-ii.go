package main

func main() {

}

func getRow(rowIndex int) []int {

	pascalsTriangle := make([][]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		pascalsTriangle[i] = make([]int, rowIndex+1)

		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				pascalsTriangle[i][j] = 1
			} else {
				pascalsTriangle[i][j] = pascalsTriangle[i-1][j-1] + pascalsTriangle[i-1][j]
			}

		}
	}
	return pascalsTriangle[rowIndex]
}
