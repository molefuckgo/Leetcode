package main

func main() {

}

func longestSubsequence(arr []int, difference int) int {
	arrLen := len(arr)
	if arrLen == 0 {
		return 0
	}
	maxLength := 1
	memoMap := make(map[int]int)
	memoMap[arr[0]] = 1
	for i := 1; i < arrLen; i++ {
		if length, ok := memoMap[arr[i]-difference]; ok {
			memoMap[arr[i]] = length + 1
		} else {
			memoMap[arr[i]] = 1
		}
		maxLength = max2(memoMap[arr[i]], maxLength)
	}
	return maxLength
}

func max2(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
