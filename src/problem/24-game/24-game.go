package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(judgePoint24([]int{4, 1, 8, 7}))
	//fmt.Println(1e-6 < 0.000001)
}

const (
	TARGET                          = 24
	EPSILON                         = 1e-6
	ADD, MULTIPLY, SUBTRACT, DIVIDE = 0, 1, 2, 3
)

func judgePoint24(nums []int) bool {
	var list []float64
	for _, num := range nums {
		list = append(list, float64(num))
	}
	return solve(list)
}

func solve(list []float64) bool {
	if len(list) == 0 {
		return false
	}
	if len(list) == 1 {
		return math.Abs(list[0]-TARGET) < EPSILON
	}
	size := len(list)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i != j {
				list2 := []float64{}
				for k := 0; k < size; k++ {
					if k != i && k != j {
						list2 = append(list2, list[k])
					}
				}
				for k := 0; k < 4; k++ {
					if k < 2 && i < j {
						continue
					}
					switch k {
					case ADD:
						list2 = append(list2, list[i]+list[j])
					case MULTIPLY:
						list2 = append(list2, list[i]*list[j])
					case SUBTRACT:
						list2 = append(list2, list[i]-list[j])
					case DIVIDE:
						if math.Abs(list[j]) < EPSILON {
							continue
						} else {
							list2 = append(list2, list[i]/list[j])
						}
					}
					if solve(list2) {
						return true
					}
					list2 = list2[:len(list2)-1]
				}
			}
		}
	}
	return false
}
