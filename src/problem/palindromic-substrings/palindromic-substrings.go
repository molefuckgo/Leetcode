package main

import "fmt"

func main() {
	fmt.Println(countSubstrings("fdsklf"))
}

func countSubstrings(s string) int {
	sLen := len(s)
	if sLen <= 1 {
		return sLen
	}
	sTemp := "#"
	for i := 0; i < len(s); i++ {
		sTemp = sTemp + string(s[i]) + "#"
	}
	//fmt.Println(sTemp)
	result := sLen
	for i := 0; i < len(sTemp); i++ {
		j := i
		k := i
		for j > 0 && k < len(sTemp)-1 {
			j--
			k++
			//fmt.Println(j, k)
			if sTemp[j] == sTemp[k] && sTemp[k] != '#' {
				result++
			} else if sTemp[k] != sTemp[j] {
				break
			}

		}

	}
	return result
}
