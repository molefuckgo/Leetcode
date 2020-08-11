package main

import "fmt"

func main() {
	node1 := Node{
		Val:       1,
		Neighbors: nil,
	}
	node2 := Node{
		Val:       2,
		Neighbors: nil,
	}
	node3 := Node{
		Val:       2,
		Neighbors: nil,
	}
	node4 := Node{
		Val:       2,
		Neighbors: nil,
	}
	node1.Neighbors = []*Node{&node2, &node4}
	node2.Neighbors = []*Node{&node1, &node3}
	node3.Neighbors = []*Node{&node1, &node2}
	node4.Neighbors = []*Node{&node1, &node3}
	fmt.Println(cloneGraph(&node1))
}

type Node struct {
	Val       int
	Neighbors []*Node
}

/*
class Solution {
	public:
		Node* used[101];           //创建一个节点（指针）数组记录每个拷贝过的节点
		Node* cloneGraph(Node* node) {
			if(!node)return node;   //如果是空指针，则返回空
			if(used[node->val])return used[node->val];  //该节点已经拷贝，直接返回该节点的指针即可
			Node* p=new Node(node->val);    //创建拷贝节点
			used[node->val]=p;             //递归会遍历每一个原有节点，然后将拷贝后的指针放入used
			vector<Node*> tp=node->neighbors;
			for(int i=0;i<tp.size();i++) //将该节点的邻接节点放入拷贝节点邻接数组
			p->neighbors.push_back(cloneGraph(tp[i]));//递归实现每一个节点的更新
			return p;           //返回拷贝后的节点
		}
};
*/

var used []*Node

func cloneGraph(node *Node) *Node {
	used = make([]*Node, 100)
	return dfsClone(node)
}

func dfsClone(node *Node) *Node {
	if used[node.Val] != nil {
		return used[node.Val]
	}
	p := &Node{
		Val: node.Val,
	}
	used[node.Val] = p
	for i := 0; i < len(node.Neighbors); i++ {
		p.Neighbors = append(p.Neighbors, dfsClone(node.Neighbors[i]))
	}

	return p
}
