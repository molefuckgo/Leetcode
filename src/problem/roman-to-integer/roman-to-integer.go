package main

import "fmt"

func main() {
	fmt.Println(romanToInt("IV"))
}

func romanToInt(s string) int {
	romanLetter2IntMap := make(map[string]int)
	romanLetter2IntMap["I"] = 1
	romanLetter2IntMap["V"] = 5
	romanLetter2IntMap["X"] = 10
	romanLetter2IntMap["L"] = 50
	romanLetter2IntMap["C"] = 100
	romanLetter2IntMap["D"] = 500
	romanLetter2IntMap["M"] = 1000
	result := 0
	lastVaule := 0
	for index, romanLetter := range s {
		if index == 0 || lastVaule > romanLetter2IntMap[string(romanLetter)] {
			result += romanLetter2IntMap[string(romanLetter)]
			lastVaule = romanLetter2IntMap[string(romanLetter)]
		} else {
			result += romanLetter2IntMap[string(romanLetter)] - 2*lastVaule
		}

	}
	return result
}
