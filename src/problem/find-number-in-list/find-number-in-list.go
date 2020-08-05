package main

import "fmt"

func main() {
	// 给定一个数字和排序好的数组（小到大），找到给定数字下标，没有返回-1
	fmt.Println(findNumberInList2(7, []int{1, 2, 3, 4, 7, 8, 9}))
}

func findNumberInList(num int, numList []int) int {
	numListLen := len(numList)

	start := 0
	end := numListLen - 1
	mid := (start + end) / 2
	for start <= end {
		if numList[mid] == num {
			return mid
		} else if numList[mid] >= num {
			end = mid - 1
			mid = (start + end) / 2
		} else {
			start = mid + 1
			mid = (start + end) / 2
		}
	}
	return -1
}

func findNumberInList2(num int, numList []int) int {
	mid := (len(numList) - 1) / 2
	if mid == 0 {
		if numList[mid] == num {
			return mid
		} else {
			return -1
		}
	}
	if numList[mid] == num {
		return mid
	} else if numList[mid] > num {
		return findNumberInList2(num, numList[:mid])
	} else {
		return findNumberInList2(num, numList[mid+1:]) + mid + 1
	}

}
