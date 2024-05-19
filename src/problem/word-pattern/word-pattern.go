package main

import (
	"fmt"
	strings "strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}

func wordPattern(pattern string, s string) bool {
	var words = strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	var hashTable1 = make(map[string]string)
	var hashTable2 = make(map[string]string)
	for i := 0; i < len(words); i++ {
		if "" == hashTable1[string([]rune(pattern)[i])] && "" == hashTable2[words[i]] {
			hashTable1[string([]rune(pattern)[i])] = words[i]
			hashTable2[words[i]] = string([]rune(pattern)[i])
		} else if hashTable1[string([]rune(pattern)[i])] != words[i] {
			return false
		}
	}
	return true
}
