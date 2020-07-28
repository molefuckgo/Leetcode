package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	str := "   -42"

	fmt.Println(myAtoi(str))
	//fmt.Println(math.MinInt32)
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

//  			' '		+/-		number		other
//	start		start	signed	in_number	end
//	signed		end		end		in_number	end
//	in_number	end		end		in_number	end
//	end			end		end 	end			end
var statusTable = [][]int{
	{0, 1, 2, 3},
	{3, 3, 2, 3},
	{3, 3, 2, 3},
	{3, 3, 3, 3},
}

const (
	start     = 0
	signed    = 1
	in_number = 2
	end       = 3

	space  = 0
	sign   = 1
	number = 2
	other  = 3
)

func myAtoi(str string) int {
	result := 0
	nowStatus := start
	positiveorNegative := true
	for _, s := range str {
		nowStatus = getStatus(nowStatus, s)
		if nowStatus == end {
			break
		} else if nowStatus == signed {
			if s == '-' {
				positiveorNegative = false
			}
		} else if nowStatus == in_number {
			result = result*10 + str2Num(string(s))
			if (result > math.MaxInt32 && positiveorNegative) || (!positiveorNegative && result < math.MinInt32) {
				if positiveorNegative {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			}
		}
	}
	if positiveorNegative {
		if result > math.MaxInt32 {
			return math.MaxInt32
		}
		return result
	} else {
		if -result <= math.MinInt32 {
			return math.MinInt32
		}
		return -result
	}

}

func getStatus(beforeStatus int, s int32) int { // 根据上次状态和现在的字符返回当前的状态
	var sType int
	if s == ' ' {
		sType = space
	} else if s == '-' || s == '+' {
		sType = sign
	} else if s >= '0' && s <= '9' {
		sType = number
	} else {
		sType = other
	}
	return statusTable[beforeStatus][sType]
}

func str2Num(str string) int {
	if intStr, err := strconv.Atoi(str); err == nil {
		return intStr
	}
	return -1
}
