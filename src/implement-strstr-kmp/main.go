package main

import "fmt"

func main() {
	txt := ""
	pat := ""
	fmt.Println(strStr(txt, pat))
}

func strStr(haystack string, needle string) int {
	lenNeedle := len(needle)
	lenHaystack := len(haystack)
	if lenNeedle == 0 {
		return 0
	} else if lenHaystack == 0 {
		return -1
	}
	next := getNext(haystack)
	j := 0
	for i := 0; i < lenHaystack; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j]
		}
		if haystack[i] == needle[j] {
			j += 1
		}
		if j == lenNeedle {
			return i - lenNeedle + 1
		}

	}
	return -1
}

func getNext(pat string) []int {
	lenPat := len(pat)
	var next = make([]int, lenPat+1)
	next[0] = -1
	j := -1

	for i := 0; i < lenPat; i++ {
		for j >= 0 && pat[i] != pat[j] {
			j = next[j]
		}
		j += 1
		next[i+1] = j
	}
	return next
}
