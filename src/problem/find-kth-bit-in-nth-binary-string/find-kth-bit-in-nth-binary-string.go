package main

import "fmt"

func main() {
	fmt.Println(string(findKthBit(20, 1048575)))
}

func findKthBit(n int, k int) byte {
	str := getNTimesString(n)
	fmt.Println(str)
	return str[k-1]
}

func getNTimesString(n int) string {
	str := "0"
	for i := 0; i < n-1; i++ {
		strReverseInverStr := invertStr(str)
		str = str + "1" + strReverseInverStr
	}
	return str
}

func invertStr(str string) string {
	newStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '0' {
			newStr += "1"
		} else {
			newStr += "0"
		}
	}
	return newStr
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// func getStringLength(n) int {
//     length := 1
//     for i := 0;i < n - 1;i++ {
//         length = length * 2 + 1
//     }
//     return length
// }
