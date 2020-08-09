package main

import (
	"fmt"
)

func main() {
	fmt.Println(countBinarySubstrings("00111000"))
}

func countBinarySubstrings(s string) int {
	sLen := len(s)
	if sLen <= 1 {
		return 0
	}
	curNum := s[0]
	beforeNumTimes := 0
	afterNumTimes := 1
	result := 0

	for i := 1; i < sLen; i++ {
		if curNum == s[i] {
			afterNumTimes++
		} else {
			curNum = s[i]
			result += min2(beforeNumTimes, afterNumTimes)
			beforeNumTimes = afterNumTimes
			afterNumTimes = 1
		}

	}
	result += min2(beforeNumTimes, afterNumTimes)
	return result

}

func min2(x, y int) int {
	if x <= y {
		return x
	}

	return y
}
