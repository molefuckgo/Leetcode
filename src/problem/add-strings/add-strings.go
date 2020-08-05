package main

import "fmt"

func main() {

	//fmt.Println(addStrings("123456", "4444"))
	fmt.Println("0")
}

func addStrings(num1 string, num2 string) string {
	num1Len := len(num1)
	num2Len := len(num2)
	if num1Len == 0 || num1[0] == '0' {
		return num2
	} else if num2Len == 0 || num2[0] == '0' {
		return num1
	}

	var maxLen, subtractLen int // 相减
	if num1Len >= num2Len {
		maxLen = num1Len
		subtractLen = num1Len - num2Len
	} else {
		maxLen = num2Len
		subtractLen = num2Len - num1Len
	}
	carry := 0
	if num1Len >= num2Len {
		for i := 0; i < subtractLen; i++ {
			num2 = "0" + num2
		}
	} else {
		for i := 0; i < subtractLen; i++ {
			num1 = "0" + num1
		}
	}
	result := ""
	var thisResult uint8
	for i := maxLen - 1; i >= 0; i-- {
		thisResult, carry = addTwoNumint32(num1[i], num2[i], carry)
		result = string(thisResult) + result
	}
	if carry == 1 {
		result = "1" + result
	}
	return result
}

func addTwoNumint32(x, y uint8, z int) (uint8, int) { // 相加后的值，是否进位
	result := x + y + uint8(z) - '0'
	var carry int
	if result >= '9'+1 {
		carry = 1
		result -= ('9' - '0' + 1)
	} else {
		carry = 0
	}
	return result, carry
}
