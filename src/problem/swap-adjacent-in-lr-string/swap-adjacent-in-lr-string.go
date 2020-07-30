package main

import "fmt"

func main() {
	start := "XXXLXXRXXR"
	end := "XXXLXXXRRX"
	//"XXXLXXXXXX"
	//"XXXLXXXXXX"
	//fmt.Println(canTransform(start, end))
	fmt.Println(canTransform(start, end))
}

func deleteX(str string, index int) string {
	if index == len(str)-1 {
		return str[0:index]
	}
	str = str[0:index] + str[index+1:]
	return str
}

func canTransform(start string, end string) bool {
	startXList := make([]int, 0)
	endtXList := make([]int, 0)
	RStartIndexList := make([]int, 0)
	REndIndexList := make([]int, 0)
	LStartIndexList := make([]int, 0)
	LEndIndexList := make([]int, 0)
	for index, letter := range start {
		if letter == 'X' {
			startXList = append(startXList, index-len(startXList))
		} else if letter == 'R' {
			RStartIndexList = append(RStartIndexList, index)
		} else { // L
			LStartIndexList = append(LStartIndexList, index)
		}
	}
	for index, letter := range end {
		if letter == 'X' {
			endtXList = append(endtXList, index-len(endtXList))
		} else if letter == 'R' {
			REndIndexList = append(REndIndexList, index)
		} else { // L
			LEndIndexList = append(LEndIndexList, index)
		}
	}
	for i := 0; i < len(startXList); i++ {
		start = deleteX(start, startXList[i])
	}
	for i := 0; i < len(endtXList); i++ {
		end = deleteX(end, endtXList[i])
	}
	if start == end {
		for i := 0; i < len(RStartIndexList); i++ {
			if RStartIndexList[i] <= REndIndexList[i] {
				continue
			}
			return false
		}
		for i := 0; i < len(LStartIndexList); i++ {
			if LStartIndexList[i] >= LEndIndexList[i] {
				continue
			}
			return false
		}
		return true
	} else {
		return false
	}

}
