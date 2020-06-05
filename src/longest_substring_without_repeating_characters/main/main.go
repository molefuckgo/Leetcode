package main

func main() {
	s := "abcabcdef"
	println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	k := -1
	res := 0
	cMap := map[string]int{}
	//sLen := len(s)
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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
