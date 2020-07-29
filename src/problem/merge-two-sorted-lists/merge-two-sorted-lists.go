package main

import "fmt"

func main() {
	l1 := ListNode{
		Val:  1,
		Next: nil,
	}
	l1.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	l1.Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}

	l2 := ListNode{
		Val:  1,
		Next: nil,
	}
	l2.Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	l2.Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}

	result := mergeTwoLists(&l1, &l2)
	resultList := make([]int, 0)
	for result != nil {
		resultList = append(resultList, result.Val)
		result = result.Next
	}
	fmt.Println(resultList)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	//l1Next := l1.Next
	//l2Next := l2.Next
	if l1 == nil && l2 != nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	result := l1
	var l1Pre *ListNode
	l1Cur := l1
	l2Cur := l2
	if l1.Val >= l2.Val {
		result = l2
	}
	for l1Cur != nil && l2Cur != nil {
		if l1Cur.Val >= l2Cur.Val {
			l2Next := l2Cur.Next
			l2Cur.Next = l1Cur
			if l1Pre != nil {
				l1Pre.Next = l2Cur
			}

			l1Pre = l2Cur
			l2Cur = l2Next
		} else if l1Cur.Next != nil {
			l1Pre = l1Cur
			l1Cur = l1Cur.Next
		} else {
			break
		}
	}
	if l2Cur != nil {
		l1Cur.Next = l2Cur
	}
	return result
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := ListNode{
		Val: -1,
	}
	var prev *ListNode
	prev = &prehead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 != nil {
		prev.Next = l1
	} else {
		prev.Next = l2
	}
	return prehead.Next
}
