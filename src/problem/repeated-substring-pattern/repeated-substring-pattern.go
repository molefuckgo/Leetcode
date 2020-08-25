package main

import "fmt"

func main() {
	s := "abaababaab"
	fmt.Println(repeatedSubstringPattern(s))
	//var a, b int
	//fmt.Scanln(&a)
	//fmt.Scanln(&b)
	//fmt.Println(a + b)
}

func repeatedSubstringPattern(s string) bool {
	sLen := len(s)
	lastLetter := s[sLen-1]
	for i := 0; i < sLen/2; i++ { // 循环找到最后一个节点
		if sLen%(i+1) != 0 {
			continue
		}
		if s[i] == lastLetter {
			flag := 0
			for j := i + 1; j < sLen-i; j += i + 1 {
				if s[0:i+1] == s[j:j+i+1] {
					flag++
					continue
				} else {
					break
				}
			}
			if flag == sLen/(i+1)-1 {
				return true
			} else {
				continue
			}
		}
	}
	return false
}

//func repeatedSubstringPattern(s string) bool {
//	sLen := len(s)
//	lastLetter := s[sLen-1]
//	for i := 0;i < sLen / 2;i++ {   // 循环找到最后一个节点
//		if s[i] == lastLetter {
//			flag := 0
//			for j := i + 1;j < sLen - i;j += i+1 {
//				if s[0:i+1] == s[j:j+i+1] {
//					flag++
//					continue
//				} else {
//					break
//				}
//			}
//			if flag == sLen / (i+1) - 1 && sLen % (i+1) == 0 {
//				return true
//			} else {
//				continue
//			}
//		}
//	}
//	return false
//}
