package main

func main() {
	l1 := new(ListNode)
	l2 := new(ListNode)
	l1.Val = 5
	l1.Next = new(ListNode)
	//l1.Next.Val = 2
	//l1.Next.Next = new(ListNode)
	//l1.Next.Next.Val = 1

	l2.Val = 5
	l2.Next = new(ListNode)
	//l2.Next.Val = 9
	//l2.Next.Next = new(ListNode)
	//l2.Next.Next.Val = 7
	listNodeResult := addTwoNumbers(l1, l2)
	for listNodeResult != nil {
		print(listNodeResult.Val)
		listNodeResult = listNodeResult.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//找长度 -> 补0 -> 记录结果的ListNode -> (是否进位, l3赋值给2, 记录相加的结果) ->
	p := l1
	q := l2
	len1 := 1
	len2 := 1
	for p.Next != nil {
		len1++
		p = p.Next
	}

	for q.Next != nil {
		len2++
		q = q.Next
	}

	if len1 > len2 {
		for i := 1; i <= len1-len2; i++ {
			q.Next = &ListNode{0, nil}
			q = q.Next
		}
	} else {
		for i := 1; i <= len2-len1; i++ {
			p.Next = &ListNode{0, nil}
			p = p.Next
		}
	}
	l3 := &ListNode{-1, nil}
	count := 0 //进位
	w := l3
	i := 0

	for l1 != nil && l2 != nil {
		i = count + l1.Val + l2.Val
		if i >= 10 {
			count = 1
		} else {
			count = 0
		}
		w.Next = &ListNode{i % 10, nil}
		l1 = l1.Next
		l2 = l2.Next
		w = w.Next
	}
	if count > 0 {
		w.Next = &ListNode{1, nil}
	}
	return l3.Next
}
