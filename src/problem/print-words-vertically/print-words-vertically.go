package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(printVertically(""))
}

func printVertically(s string) []string {
	var words = strings.Split(s, " ")
	var result []string
	var max int
	for _, word := range words {
		if max < len(word) {
			max = len(word)
		}
	}

	for i := 0; i < max; i++ {
		result = append(result, "")
		for _, word := range words {
			if i < len(word) {
				result[i] += string(word[i])
			} else {
				result[i] += " "
			}
		}
		result[i] = strings.TrimRight(result[i], " ")
	}
	return result
}
