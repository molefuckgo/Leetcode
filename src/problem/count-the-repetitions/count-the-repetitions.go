package main

import "fmt"

func main() {
	fmt.Println(getMaxRepetitions("abaacdbac", 3, "adcbd", 1))
}

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if n1 == 0 {
		return 0
	}
	s1Cnt, index, s2Cnt := 0, 0, 0
	recall := make(map[int][]int)
	var preLoop, inLoop []int
	for true {
		// 多遍历一个s1，看看能不能找到循环节
		s1Cnt++
		for _, ch := range s1 {
			if ch == int32(s2[index]) {
				index++
				if index == len(s2) {
					s2Cnt, index = s2Cnt+1, 0
				}
			}

		}
		// 还没找到循环节， 所有的s1就用完了
		if s1Cnt == n1 {
			return s2Cnt / n2
		}
		if _, ok := recall[index]; ok {
			recallIndex := recall[index]
			s1CntPrime, s2CntPrime := recallIndex[0], recallIndex[1]
			// 前 s1Cnt' 个 s1 包含了 s2Cnt' 个s2
			preLoop = []int{s1CntPrime, s2CntPrime}
			// 以后每(s1Cnt - s1Cnt')个s1包含了（s2Cnt - s2Cnt'）个s2
			inLoop = []int{s1Cnt - s1CntPrime, s2Cnt - s2CntPrime}
			break
		} else {
			recall[index] = []int{s1Cnt, s2Cnt}
		}

	}
	// ans存储的是s1包含的s2的数量，考虑的之前的pre_loop和in_loop
	ans := preLoop[1] + (n1-preLoop[0])/inLoop[0]*inLoop[1]
	// s1的末尾还剩下一些s1，我们暴力进行匹配
	rest := (n1 - preLoop[0]) % inLoop[0]
	for i := 0; i < rest; i++ {
		for _, ch := range s1 {
			if ch == int32(s2[index]) {
				index++
				if index == len(s2) {
					ans, index = ans+1, 0
				}
			}

		}
	}
	// s1包含ans个s2， 那么就包含ans / n2个s2
	return ans / n2
}
