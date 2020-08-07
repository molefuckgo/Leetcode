package main

import "fmt"

func main() {
	capacity := 2
	obj := Constructor(capacity)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(1))
	obj.Put(4, 4)
}

//type LinkNode struct {
//	Key, Val int
//	Next, Prev *LinkNode
//}
//
//
//
//type LRUCache struct {
//	m map[int]*LinkNode
//	capacity int
//	head, tail *LinkNode
//}
//
//
//func Constructor(capacity int) LRUCache {
//	head, tail := &LinkNode{-1, -1, nil, nil}, &LinkNode{-1, -1, nil, nil}
//	head.Next = tail
//	tail.Prev = head
//	cache := LRUCache{make(map[int]*LinkNode), capacity, head, tail}
//	return cache
//}
//
//func (this *LRUCache) Add2Head(node * LinkNode) { // 向头部添加节点（注意真实的头部是维护的key， val都是-1的节点，同样尾节点也是）
//	node.Prev = this.head
//	node.Next = this.head.Next
//	this.head.Next = node
//	node.Next.Prev = node
//}
//
//func (this *LRUCache) RemoveNode(node *LinkNode) {
//	node.Prev.Next = node.Next
//	node.Next.Prev = node.Prev
//}
//
//func (this *LRUCache) Move2Head(node *LinkNode) {
//	this.RemoveNode(node)
//	this.Add2Head(node)
//}
//
//
//func (this *LRUCache) Get(key int) int {
//	m := this.m
//	if node, ok := m[key]; ok {
//		this.Move2Head(node)
//		return node.Val
//	} else {
//		return -1
//	}
//}
//
//
//func (this *LRUCache) Put(key int, val int)  {
//	m := this.m
//	if node, ok := m[key]; ok {
//		node.Val = val
//		this.Move2Head(node)
//	} else {
//		node = &LinkNode{
//			Key:  key,
//			Val:  val,
//			Next: nil,
//			Prev: nil,
//		}
//		if len(m) >= this.capacity {
//			delete(m, this.tail.Prev.Key)
//			this.RemoveNode(this.tail.Prev)
//		}
//		m[key] = node
//		this.Add2Head(node)
//	}
//}

type LinkNode struct {
	Key, Val   int
	Prev, Next *LinkNode
}
type LRUCache struct {
	m          map[int]*LinkNode
	capacity   int
	head, tail *LinkNode
}

func Constructor(capacity int) LRUCache {
	head := &LinkNode{
		Key:  -1,
		Val:  -1,
		Prev: nil,
		Next: nil,
	}

	tail := &LinkNode{
		Key:  -1,
		Val:  -1,
		Prev: nil,
		Next: nil,
	}
	head.Next = tail
	tail.Prev = head
	m := make(map[int]*LinkNode)
	return LRUCache{
		m:        m,
		capacity: capacity,
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Add2Head(node *LinkNode) {
	node.Next = this.head.Next
	node.Prev = this.head
	this.head.Next = node
	node.Next.Prev = node
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) Move2Head(node *LinkNode) {
	this.RemoveNode(node)
	this.Add2Head(node)
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.Move2Head(node)
		return node.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		node.Val = value
		this.Move2Head(node)
	} else {
		node := &LinkNode{
			Key:  key,
			Val:  value,
			Prev: nil,
			Next: nil,
		}
		if len(this.m) == this.capacity {
			delete(this.m, this.tail.Prev.Key)
			this.RemoveNode(node)
		}
		this.m[key] = node
		this.Add2Head(node)

	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
