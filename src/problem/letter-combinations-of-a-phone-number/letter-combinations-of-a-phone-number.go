package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("667"))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	num2letter := make(map[uint8][]string)
	num2letter[uint8('2')] = []string{"a", "b", "c"}
	num2letter[uint8('3')] = []string{"d", "e", "f"}
	num2letter[uint8('4')] = []string{"g", "h", "i"}
	num2letter[uint8('5')] = []string{"j", "k", "l"}
	num2letter[uint8('6')] = []string{"m", "n", "o"}
	num2letter[uint8('7')] = []string{"p", "q", "r", "s"}
	num2letter[uint8('8')] = []string{"t", "u", "v"}
	num2letter[uint8('9')] = []string{"w", "x", "y", "z"}
	result := make([]string, 0)
	for i := 0; i < len(num2letter[digits[0]]); i++ {
		result = append(result, num2letter[digits[0]][i])
	}
	for i := 1; i < len(digits); i++ {
		result = dfs(result, num2letter[digits[i]])
	}

	return result

}

func dfs(before, after []string) []string {
	result := make([]string, len(before)*len(after))

	if len(after) == 3 {
		for i := 0; i < len(result); i += 3 {
			result[i] += before[i/3] + after[0]
			result[i+1] += before[i/3] + after[1]
			result[i+2] += before[i/3] + after[2]
		}
	} else {
		for i := 0; i < len(result); i += 4 {
			result[i] += before[i/4] + after[0]
			result[i+1] += before[i/4] + after[1]
			result[i+2] += before[i/4] + after[2]
			result[i+3] += before[i/4] + after[3]
		}
	}

	return result
}
