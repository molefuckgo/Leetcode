package main

import (
	"fmt"
	"math"
)

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())

}

//type LinkList struct {
//	Val int
//	Next *LinkList
//}
type MinStack struct {
	dataList []int
	minList  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {

	minList := []int{math.MinInt64, math.MaxInt64}
	return MinStack{
		dataList: []int{},
		minList:  minList,
	}
}

func (this *MinStack) Push(x int) {

	this.dataList = append(this.dataList, x)
	start := 0
	end := len(this.minList) - 1
	mid := (start + end) / 2
	for start <= end {
		if this.minList[mid] <= x && this.minList[mid+1] >= x {
			this.minList = append(this.minList[:mid+1], append([]int{x}, this.minList[mid+1:]...)...)
			break
		} else {
			if this.minList[mid] > x {
				end = mid - 1
			} else {
				start = mid + 1
			}
			mid = (start + end) / 2
		}
	}
}

func (this *MinStack) Pop() {
	start := 0
	end := len(this.minList) - 1
	mid := (start + end) / 2
	for start <= end {
		if this.minList[mid] == this.dataList[len(this.dataList)-1] {
			this.minList = append(this.minList[:mid], this.minList[mid+1:]...)
			break
		} else {
			if this.minList[mid] > this.dataList[len(this.dataList)-1] {
				end = mid - 1
			} else {
				start = mid + 1
			}
			mid = (start + end) / 2
		}
	}
	this.dataList = this.dataList[:len(this.dataList)-1]
}

func (this *MinStack) Top() int {
	return this.dataList[len(this.dataList)-1]
}

func (this *MinStack) GetMin() int {
	return this.minList[1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
