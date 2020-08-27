package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findItinerary([][]string{
		{"EZE","TIA"},
		{"EZE","HBA"},
		{"AXA","TIA"},
		{"JFK","AXA"},
		{"ANU","JFK"},
		{"ADL","ANU"},
		{"TIA","AUA"},
		{"ANU","AUA"},
		{"ADL","EZE"},
		{"ADL","EZE"},
		{"EZE","ADL"},
		{"AXA","EZE"},
		{"AUA","AXA"},
		{"JFK","AXA"},
		{"AXA","AUA"},
		{"AUA","ADL"},
		{"ANU","EZE"},
		{"TIA","ADL"},
		{"EZE","ANU"},
		{"AUA","ANU"},
	}))


}

func findItinerary(tickets [][]string) []string {
	var (
		m  = map[string][]string{}
		res []string
	)

	for _, ticket := range tickets {	// 构造map
		src, dst := ticket[0], ticket[1]
		m[src] = append(m[src], dst)
	}
	for key := range m {	// 排序
		sort.Strings(m[key])
	}

	var dfs func(curr string)
	dfs = func(curr string) {
		for {
			if v, ok := m[curr]; !ok || len(v) == 0 {
				break
			}
			tmp := m[curr][0]
			m[curr] = m[curr][1:]
			dfs(tmp)
		}
		res = append(res, curr)
	}

	dfs("JFK")
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res) - 1 - i] = res[len(res) - 1 - i], res[i]
	}
	return res
}


//func findItinerary(tickets [][]string) []string {
//	memoTicketsMap := make(map[string][]string)
//	for i := 0;i < len(tickets);i++ {
//		memoTicketsMap[tickets[i][0]] = append(memoTicketsMap[tickets[i][0]], tickets[i][1])
//	}
//	result := []string{"JFK"}
//	nowKey := "JFK"
//	minIndex := 0
//	for len(memoTicketsMap) != 0 {
//		nowKeyTemp := nowKey
//		nowKey, minIndex = getMinStringAndIndex(memoTicketsMap[nowKey], &memoTicketsMap)
//		result = append(result, nowKey)
//		memoTicketsMap[nowKeyTemp] = deleteArrayIndex(memoTicketsMap[nowKeyTemp], minIndex)
//		if len(memoTicketsMap[nowKeyTemp]) == 0 {
//			delete(memoTicketsMap, nowKeyTemp)
//		}
//		fmt.Println(result)
//	}
//	return result
////	[ JFK, AXA, AUA , ADL , ANU, AUA, ANU, EZE, ADL, EZE, ANU, JFK, AXA, EZE, TIA, AUA, AXA, TIA , ADL , EZE , HBA ]
////	[ JFK, AXA, AUA , ADL , ANU, AUA, ANU, EZE, ADL, EZE, ANU, JFK, AXA, EZE, TIA, ADL, EZE, HBA]
//
//}
//
//func getMinStringAndIndex(tickets []string, memoTicketsMap *map[string][]string) (string, int) {
//	minString := tickets[0]
//	var minIndex int
//	for i := 1;i < len(tickets);i++ {
//		if len((*memoTicketsMap)[minString]) == 0 || (minString > tickets[i] && len((*memoTicketsMap)[tickets[i]]) != 0) {
//			minIndex = i
//			minString = tickets[i]
//		}
//	}
//	return minString, minIndex
////	&& len((*memoTicketsMap)[tickets[i]]) != 0
//}
//
//func deleteArrayIndex(array []string, index int) []string {
//	array = append(array[0:index], array[index+1:]...)
//	return array
//}
