package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abcbbcb"))
}

func longestPalindrome(s string) string {
	lenS := len(s)
	if lenS < 2 {
		return s
	}
	t := "#"
	for i := 0; i < lenS; i++ {
		t += string(s[i]) + "#"
	}
	lenT := 2*lenS + 1
	p := make([]int, lenT)
	// 双指针，它们是一一对应的，必须同时更新
	maxRight := 0
	center := 0

	maxLen := 1 // 当前遍历的中心最大扩散步数，其值等于原始最富川的最长回文字串的长度
	start := 1

	for i := 0; i < lenT; i++ {
		if i < maxRight {
			mirror := 2*center - i
			p[i] = min(maxRight-i, p[mirror])
		}
		// 下一次尝试扩散的左右七点，能扩散呢的步数直接加到p[i]中
		left := i - (1 + p[i])
		right := i + (1 + p[i])

		for left > 0 && right < lenT && t[left] == t[right] {
			p[i]++
			left--
			right++
		}
		if i+p[i] > maxRight {
			// 同时更新maxRight和center
			maxRight = i + p[i]
			center = i
		}

		if p[i] > maxLen {
			// 记录最长回文子串的长度和相应它在原始字符串中的起点
			maxLen = p[i]
			start = (i - maxLen) / 2
		}
	}
	return s[start : start+maxLen]
}

func min(x int, y int) int {
	if x >= y {
		return y
	}
	return x
}
func longestPalindrome2(s string) string {
	lenS := len(s)
	if lenS < 2 {
		return s
	}
	t := "#"
	for i := 0; i < lenS; i++ {
		t += string(s[i]) + "#"
	}
	lenT := 2*lenS + 1
	maxLen := -1
	start := 0
	// p := make([]int, lenT)
	for i := 0; i < lenT; i++ {
		curStep := centerSpread(t, i, lenT)
		if curStep > maxLen {
			maxLen = curStep
			start = (i - maxLen) / 2
		}
	}
	return s[start : start+maxLen]
}

func centerSpread(s string, center int, lenT int) int {
	i := center - 1
	j := center + 1
	step := 0
	for i > 0 && j < lenT && s[i] == s[j] {
		step++
		i--
		j++
	}
	return step
}
