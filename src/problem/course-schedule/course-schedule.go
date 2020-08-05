package main

import (
	"fmt"
)

func main() {
	fmt.Println(canFinish(6, [][]int{
		{0, 1},
		{2, 3},
		{3, 4},
		{4, 2},
	}))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	admission := make([]int, numCourses)            //入度
	admissionStartWith := make([][]int, numCourses) // 学完后有影响对列表
	zeroAdmission := make([]int, 0)                 // 入度为0的课程列表
	for _, courseInfo := range prerequisites {
		admission[courseInfo[0]]++
		admissionStartWith[courseInfo[1]] = append(admissionStartWith[courseInfo[1]], courseInfo[0])
	}
	for i := 0; i < numCourses; i++ {
		if admission[i] == 0 {
			zeroAdmission = append(zeroAdmission, i)
		}
	}
	result := 0
	for len(zeroAdmission) != 0 {
		nowCourse := zeroAdmission[0]
		zeroAdmission = zeroAdmission[1:]
		result++
		// 减法
		//for i := 0;i < len(admission);i++ {
		for _, shouldSub := range admissionStartWith[nowCourse] {
			admission[shouldSub]--
			if admission[shouldSub] == 0 {
				zeroAdmission = append(zeroAdmission, shouldSub)
			}
			//}

		}

	}
	return result == numCourses
}
