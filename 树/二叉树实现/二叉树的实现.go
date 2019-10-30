package main

import (
	"fmt"

)

type Node struct {
	data  int
	left  *Node
	right *Node
}

//先序遍历
func preorder(node *Node) {
	if node != nil {
		fmt.Println(node.data)
	preorder(node.left)
		preorder(node.right)
	}
}

//中序遍历
func inorder(node *Node) {
	if node != nil {
		inorder(node.left)
		fmt.Println(node.data)
		inorder(node.right)
	}
}

//后序
func postorder(node *Node) {
	if node != nil {
		postorder(node.left)
		postorder(node.right)
		fmt.Println(node.data)
	}

}

func main() {
	n1 := Node{}
	n2 := Node{}
	n3 := Node{}
	n4 := Node{}
	n1.data = 5
	n2.data = 6
	n3.data = 7
	n4.data = 8
	n1.left = &n2
	n1.right = &n3
	n2.left = &n4
	n2.right = nil
	n3.right = nil
	n3.left = nil
	n4.left = nil
	n4.right = nil

	//preorder(&n1)
	//inorder(&n1)
	postorder(&n1)
}
