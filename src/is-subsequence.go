package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("acb", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	if sLen == 0 {
		return true
	} else if tLen == 0 || sLen > tLen {
		return false
	}

	tIndex := 0
	for sIndex := 0; sIndex < sLen; {
		for tIndex < tLen {
			if s[sIndex] == t[tIndex] {
				if sIndex == sLen-1 {
					return true
				}
				if tIndex == tLen-1 {
					return false
				}
				sIndex++
				tIndex++
			} else {
				if tIndex == tLen-1 {
					return false
				}
				tIndex++
			}
		}
	}
	return false
}
