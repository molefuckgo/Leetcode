package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(multiply("999", "999"))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	num1Len, num2Len := len(num1), len(num2)
	var result string // 用于存储结果

	resultArr := make([][]int16, 0)

	for i := 0; i < num2Len; i++ {
		oneTimeResult := make([]int16, num1Len+num2Len-1)
		for j := 0; j < num1Len; j++ {
			oneTimeResult[num1Len+num2Len-2-j-i] = int16((num2[i] - '0') * (num1[j] - '0'))
		}
		resultArr = append(resultArr, oneTimeResult)
	}
	var carry int16
	carry = 0
	// 进位
	for i := 0; i <= num1Len+num2Len-2; i++ {
		var columnAddresult int16
		columnAddresult = carry
		for j := 0; j < len(resultArr); j++ {
			columnAddresult += resultArr[j][i]
		}

		result = strconv.Itoa(int((columnAddresult % 10))) + result
		carry = columnAddresult / 10
	}
	if carry > 0 {
		return strconv.Itoa(int(carry)) + result
	}
	return result
}
