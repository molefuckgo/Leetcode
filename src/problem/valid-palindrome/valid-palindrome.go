package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println('a')	// 97
	//fmt.Println('z')	// 122
	//fmt.Println('A')	// 65
	//fmt.Println('Z')	//90
	fmt.Println(isPalindrome("0P"))
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	sLen := len(s)
	leftPtr := 0
	rightPtr := sLen - 1
	for leftPtr < rightPtr {
		leftIsLetter := isLetter(s[leftPtr])
		leftIsNumber := isNumber(s[leftPtr])
		rightIsLetter := isLetter(s[rightPtr])
		rightIsNumber := isNumber(s[rightPtr])
		if (leftIsLetter && rightIsNumber) || (leftIsNumber && rightIsLetter) {
			return false
		} else if leftIsLetter && rightIsLetter {
			if judgeLettersIsEqual(s[leftPtr], s[rightPtr]) { // 判断后相等
				leftPtr++
				rightPtr--
			} else {
				return false
			}
		} else if leftIsNumber && rightIsNumber {
			if s[leftPtr] == s[rightPtr] {
				leftPtr++
				rightPtr--
			} else {
				return false
			}
		} else if !leftIsLetter && !leftIsNumber {
			leftPtr++
			continue
		} else {
			rightPtr--
			continue
		}
	}
	return true
}

func isLetter(letter uint8) bool {
	if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
		return true
	} else {
		return false
	}
}

func isNumber(letter uint8) bool {
	return letter >= '0' && letter <= '9'
}

func judgeLettersIsEqual(letter1, letter2 uint8) bool { // 判断两个字符是否相等，忽略大小写
	if letter1 == letter2 || math.Abs(float64(letter1)-float64(letter2)) == 32 {
		return true
	} else {
		return false
	}

}
