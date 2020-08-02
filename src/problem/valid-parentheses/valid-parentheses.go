package main

import "fmt"

func main() {
	fmt.Println(isValid("){"))

}

var (
	space         = 0
	left          = 1
	right         = 2
	left2RightMap = make(map[int32]int32)
)

func isValid(s string) bool {
	lenS := len(s)
	if lenS == 0 {
		return true
	} else if lenS%2 == 1 {
		return false
	}
	left2RightMap['{'] = '}'
	left2RightMap['('] = ')'
	left2RightMap['['] = ']'
	signL := make([]int32, 0)
	for _, sign := range s {
		whickSign := judgeLeftOrRightOrSpace(sign)
		if whickSign == space {
			continue
		} else if whickSign == left {
			signL = append(signL, sign)
		} else {
			if len(signL) != 0 && sign == left2RightMap[signL[int32(len(signL)-1)]] {
				signL = signL[:len(signL)-1]
			} else {
				return false
			}
		}
	}
	if len(signL) != 0 {
		return false
	}
	return true
}

func judgeLeftOrRightOrSpace(s int32) int {
	if s == ']' || s == '}' || s == ')' {
		return right
	} else if s == ' ' {
		return space
	} else {
		return left
	}
}
