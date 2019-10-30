package main

import (
	"fmt"
)

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	//判断是不是添加第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //构成一个环形
		fmt.Println(newCatNode, "加入到环形的链表")
		return
	}

	//定义一个临时变量，帮忙，找到环形的最后结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	//加入到链表中
	temp.next = newCatNode
	newCatNode.next = head

}

//输出环形链表
func ListCircle(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空空的环形链表")
		return
	}
	for {
		fmt.Printf("猫的信息=[id=%d]", temp.no)
		if temp.next == head {
			break
		}
		temp = temp.next
	}

}

func DelCatNode(head *CatNode, id int) {
	temp := head
	helper := head
	//空链表
	if head.next == nil {
		fmt.Println("无法删除")
		return
	}
	//if only 1 node
	if temp.next == head {
		temp.next = nil
	}
	flag := true
	//2 以上个node
	for {
		if temp.next == head {
			break
		}
		if temp.no == id {
			//找到了，
			helper.next = temp.next

		}
		temp = temp.next
		helper = helper.next
	}

	//这里还要比较一次
	if flag {
		if temp.no == id {
			helper.next = temp.next
		}
	}
}

func main() {
	head := &CatNode{}

	//创建一只猫
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}

	InsertCatNode(head, cat1)
	ListCircle(head)
}
