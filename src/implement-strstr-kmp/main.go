package main

import (
	"fmt"
)

func main() {
	// txt := ""
	needle := "abababaab"
	haystack := "cbabcabababaabab"
	fmt.Println(strStr(haystack, needle))
}

// func strStr(haystack string, needle string) int {
// 	lenNeedle := len(needle)
// 	lenHaystack := len(haystack)
// 	if lenNeedle == 0 {
// 		return 0
// 	} else if lenHaystack == 0 {
// 		return -1
// 	}
// 	next := getNext(haystack)
// 	j := 0
// 	for i := 0; i < lenHaystack; i++ {
// 		for j > 0 && haystack[i] != needle[j] {
// 			j = next[j]
// 		}
// 		if haystack[i] == needle[j] {
// 			j += 1
// 		}
// 		if j == lenNeedle {
// 			return i - lenNeedle + 1
// 		}
//
//
// 	}
// 	return -1
// }
//
// func getNext(pat string) []int {
// 	lenPat := len(pat)
// 	var next = make([]int, lenPat+1)
// 	next[0] = -1
// 	j := -1
//
// 	for i := 0; i < lenPat; i++ {
// 		for j >= 0 && pat[i] != pat[j] {
// 			j = next[j]
// 		}
// 		j += 1
// 		next[i+1] = j
// 	}
// 	return next
// }

// func strStr(haystack string, needle string) int {
// 	lenNeedle := len(needle)
// 	lenHaystack := len(haystack)
// 	if lenNeedle == 0 {
// 		return 0
// 	} else if lenHaystack == 0 {
// 		return -1
// 	}
// 	return 0
// }
//
// func generateNext(needle string, lenNeedle int) []int {
//
// 	nextArray := make([]int, lenNeedle)
// 	nextArray[0] = -1
// 	k := -1
// 	for q := 2; q < lenNeedle; q++ {
// 		for ;k >= 0 && needle[k + 1] != needle[q];{
// 			k = nextArray[k]
// 		}
// 		if needle[k + 1] == needle[q] {
// 			k++
// 		}
// 		nextArray[q] = k
// 	}
// 	return nextArray
// }

func strStr(haystack string, needle string) int {
	lenNeedle := len(needle)
	lenHaystack := len(haystack)
	if lenNeedle == 0 {
		return 0
	}
	if len(haystack) < lenNeedle {
		return -1
	}
	next := nextArr1(needle, lenNeedle)
	fmt.Println("next:", next)
	i := 0
	j := 0 // 当前位置(i)和needle的j位置比较
	length := 0
	for i < lenHaystack {
		if haystack[i] != needle[j] {
			if j == 0 { // 到头还是无法匹配
				i += 1
			} else {
				j = next[j-1] // 求出的next数组让当前位置和next[j - 1]去比较
			}
			length = j // 已经比较过的长度
		} else {
			length += 1
			i += 1
			j += 1
		}

		if length == lenNeedle {
			return i - length
		}
	}
	return -1 // 未找到
}

func nextArr(needle string, lenNeedle int) []int {
	next := make([]int, lenNeedle)

	for i := 1; i < lenNeedle; i++ {
		preNextVal := next[i-1]
		// fmt.Println("i:", i)
		// fmt.Println("nedle: a b b b a a b b b a c b a a b b b")
		// fmt.Println("index: 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6")
		// fmt.Println("next:", next)
		// fmt.Println("preNextVal:", preNextVal)
		for preNextVal > 0 && needle[i] != needle[preNextVal] { // 不相等循环找前面和现在（i）相等的找到后next下标就是现在应该的下标
			// fmt.Println("preNextVal > 0 && needle[i] != needle[preNextVal]")
			// fmt.Printf("next[preNextVal - 1]: %v\n", next[preNextVal - 1])
			preNextVal = next[preNextVal-1]
			// fmt.Println("preNextVal = next[preNextVal - 1]")
		}
		if needle[i] == needle[preNextVal] { // 相等让比较指针后移
			// fmt.Println("needle[i] == needle[preNextVal]")
			preNextVal += 1
			// fmt.Println("preNextVal += 1")
		} else {
			// fmt.Println("needle[i] != needle[preNextVal], preNextVal没+1")
		}
		next[i] = preNextVal //
		// fmt.Println("next[i] = preNextVal")
		// fmt.Println("next:", next)
		// fmt.Println("------------------------------------------------")
	}
	return next
}

func nextArr1(needle string, lenNeedle int) []int {
	next := make([]int, lenNeedle)
	for i := 1; i < lenNeedle; i++ {
		preNextVal := next[i-1]
		for preNextVal > 0 && needle[i] != needle[preNextVal] {
			preNextVal = next[preNextVal-1]

		}
		if needle[i] == needle[preNextVal] {
			preNextVal += 1
		}
		next[i] = preNextVal
	}
	return next
}
