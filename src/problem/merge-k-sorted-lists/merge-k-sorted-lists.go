package main

import "container/heap"

func main() {
	// 输入的二维数组
	input := [][]int{
		{},
		{1, 3, 4},
	}

	// 初始化链表数组
	var lists []*ListNode

	// 将二维数组转换为链表
	for _, arr := range input {
		// 使用一个哨兵节点简化链表构造
		dummy := &ListNode{}
		current := dummy
		for _, val := range arr {
			current.Next = &ListNode{Val: val}
			current = current.Next
		}
		// 将构造好的链表（跳过哨兵节点）添加到列表
		lists = append(lists, dummy.Next)
	}

	mergeKLists(lists)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码不保证正确性，仅供参考。如有疑惑，可以参照我写的 java 代码对比查看。

//Definition for singly-linked list.

/*func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	// 虚拟头节点
	dummy := &ListNode{Val: -1}
	p := dummy
	// 优先队列,最小堆, 用golang的heap
	pq := make(Queue, len(lists))
	for i, head := range lists {
		if head != nil {
			pq[i] = head
		}
	}
	heap.Init(&pq)

	for pq.Len() != 0 {
		// 获取最小节点，接到结果链表中
		node := heap.Pop(&pq).(*ListNode)
		p.Next = node
		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
		// p 指针不断前进
		p = p.Next
	}
	return dummy.Next
}*/

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var dummy = &ListNode{
		Val:  -1,
		Next: nil,
	}
	var q = make(Queue, 0)
	var p = dummy
	for _, head := range lists {
		if head != nil {
			q = append(q, head)
		}
	}
	heap.Init(&q)
	for q.Len() > 0 {
		var node = heap.Pop(&q).(*ListNode)
		p.Next = node
		if node.Next != nil {
			heap.Push(&q, node.Next)
		}
		p = p.Next
	}
	return dummy.Next
}

type Queue []*ListNode

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	return q[i].Val < q[j].Val
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x.(*ListNode))
}

func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}
