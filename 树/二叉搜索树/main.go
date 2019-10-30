package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func insert(tree *Tree, value int) {
	node := new(Node)
	node.data = value
	node.left = nil
	node.right = nil
	if tree.root == nil {
		tree.root = node
	} else {
		temp := tree.root
		for temp != nil {
			if value < temp.data {
				if temp.left == nil {
					temp.left = node
					return
				} else {
					temp = temp.left
				}
			} else {
				if temp.right == nil {
					temp.right = node
					return
				} else {
					temp = temp.right
				}
			}
		}
	}

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

func getHeight(node *Node) int {
	if node == nil {
		return 0
	} else {
		leftHeight := getHeight(node.left)
		rightHeight := getHeight(node.right)
		max := leftHeight
		if rightHeight > max {
			max = rightHeight
		}
		return max + 1
	}
}

func getMax(node *Node) int {
	if node == nil {
		return -1
	} else {
		m1 := getMax(node.left)
		m2 := getMax(node.right)
		m3 := node.data
		max := m1
		if m2 > max {
			max = m2
		}
		if m3 > max {
			max = m3
		}
		return max
	}

}

func main() {
	arr := []int{6, 3, 8, 2, 5, 7, 1}
	//arr:=[]int{2,8,9,3,4,6,1,5,7}
	var tree Tree
	tree.root = nil
	for i := 0; i < len(arr); i++ {
		insert(&tree, arr[i])
	}
	//preorder(tree.root)
	inorder(tree.root)
	//postorder(tree.root)

	res := getHeight(tree.root)
	fmt.Println("height=", res)
	mum := getMax(tree.root)
	fmt.Println("max=", mum)

}
