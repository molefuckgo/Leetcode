package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	//str := "  0123 4"

	//fmt.Println(myAtoi(str))
	fmt.Println(math.MinInt32)
	//	num:48 - 57
}

// https://leetcode-cn.com/problems/string-to-integer-atoi/

func myAtoi2(str string) int {
	result := 0
	sliIndex := 0
	str = strings.Trim(str, " ")
	for _, s := range str {
		if s == '0' {
			sliIndex++
		} else {
			break
		}
	}

	//if sliIndex > 0 {
	//	str = str[sliIndex:]
	//	if len(str) == 0 {
	//		return 0
	//	}
	//	if (str[0] == '+' || str[0] == '-') && sliIndex > 0 {
	//		return 0
	//	}
	//}
	if len(str) == 0 || (str[0] != '-' && str[0] != '+') && (str[0] < '0' || str[0] > '9') {
		return result
	}
	var positiveorNegative bool
	if str[0] == '-' { // 记录正负
		positiveorNegative = false
		str = str[1:]
	} else if str[0] == '+' {
		positiveorNegative = true
		str = str[1:]
	} else {
		positiveorNegative = true
	}
	sliIndex = 0
	for _, s := range str {
		if s == '0' {
			sliIndex++
		} else {
			break
		}
	}
	if sliIndex > 0 {
		str = str[sliIndex:]
	}

	for forloop, num := range str {
		if num >= '0' && num <= '9' {
			if forloop > 9 {
				if positiveorNegative {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			}
			result = result*10 + str2Num(string(num))
		} else {
			break
		}
	}
	if positiveorNegative {
		if result <= math.MaxInt32 {
			return result
		}
		return math.MaxInt32
	} else {
		if result <= math.MaxInt32 {
			return -result
		}
		return math.MinInt32
	}

}

func str2Num(str string) int {
	if intStr, err := strconv.Atoi(str); err == nil {
		return intStr
	}
	return -1
}

//  			' '		+/-		number		other
//	start		start	signed	in_number	end
//	signed		end		end		in_number	end
//	in_number	end		end		in_number	end
//	end			end		end 	end			end
