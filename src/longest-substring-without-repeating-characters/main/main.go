package main

func main() {
	s := "abcabcdef"
	println(lengthOfLongestSubstring(s))
}

// func lengthOfLongestSubstring(s string) int {
// 	k := -1
// 	res := 0
// 	cMap := map[string]int{}
// 	//sLen := len(s)
// 	for i, c := range s {
// 		_, keyInMap := cMap[string(c)]
// 		if keyInMap && cMap[string(c)] > k {
// 			k = cMap[string(c)]
// 			cMap[string(c)] = i
// 		} else {
// 			cMap[string(c)] = i
// 			res = max(res, i-k)
// 		}
// 	}
// 	return res
// }

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring(s string) int {
	// 定义变量用于记录无重复开头下标 k
	// 结果记录res
	// for {
	// 		如果c在map中存在
	// 			无重复记录开头k = cMap[string(c)]
	// 			cMap[string(c)] 记录当前循环下标
	// 		如果不是
	// 			cMap[string(c)] 记录当前循环下标
	// 			res = max(i - k, res)
	//
	// }
	//
	k := -1
	res := 0
	cMap := map[string]int{}
	for i, c := range s {
		_, keyInMap := cMap[string(c)]
		if keyInMap && cMap[string(c)] > k {
			k = cMap[string(c)]
			cMap[string(c)] = i
		} else {
			cMap[string(c)] = i
			res = max(res, i-k)
		}
	}
	return res
}
