package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(groupAnagrams([]string{"tin", "ram", "zip", "cry", "pus", "jon", "zip", "pyx"}))
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{{""}}
	}
	var recordMap = make(map[string][]string)
	for _, one_word := range strs {
		var recordArray = make([]int, 26)
		for _, oneLetter := range one_word {
			recordArray[oneLetter-'a'] += 1
		}
		recordMap[getRecordString(recordArray)] = append(recordMap[getRecordString(recordArray)], one_word)
	}
	var res [][]string
	for _, record := range recordMap {
		res = append(res, record)
	}
	return res
}

func getRecordString(recordArray []int) string {
	var recordString = ""
	for _, oneRecord := range recordArray {
		recordString += "#" + strconv.Itoa(oneRecord)
	}
	return recordString
}
