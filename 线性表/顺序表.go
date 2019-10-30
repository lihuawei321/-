package main

import (
	"errors"
	"fmt"
)

type SequenList struct {
	size   int
	length int
	data   [10]int
}

func NewSequenList() *SequenList {
	var arr [10]int
	return &SequenList{
		size:   10,
		length: 10,
		data:   arr,
	}

}

//打印线性表
func (sl *SequenList) Show() {
	fmt.Println(sl)

}

//插入元素：从末尾append一个数据
func (sl *SequenList) Append(data int) error {
	if sl.IsFull() {
		return errors.New("SequeList overflow")
	}
	sl.data[sl.length] = data
	sl.length++
	return nil
}

//插入元素，任意位置插入元素
func (sl *SequenList) Insert(index int, data int) error {
	if sl.IsFull() {
		return errors.New("SequeList overflow")
	}
	if index < 0 || index > sl.length {
		return errors.New("index overflow")
	}
	//这里如果按照正序循环则书写极其麻烦,从最后一位开始往后移动很简便
	for i := sl.length; i >= index; i-- {
		if i == index {
			sl.data[i] = data
			break
		}
		sl.data[i] = sl.data[i-1]
	}
	sl.length++
	return nil
}

// 删除元素：从末尾pop一个数据
func (sl *SequenList) Pop() (int, error) {

	if sl.Empty() {
		return 0, errors.New("SequenList is empty")
	}

	e := sl.data[sl.length-1]
	sl.data[sl.length-1] = 0
	sl.length--
	return e, nil
}

//判断顺序表是否已满
func (sl *SequenList) IsFull() bool {
	if sl.length == sl.size {
		return true
	}
	return false

}

//判断顺序表是否为空
func (sl *SequenList) Empty() bool {
	if sl.length == 0 {
		return true
	}
	return false
}

//获取线性表容量
func (sl *SequenList) Size() int {
	return sl.size

}
func main() {
	s1 := SequenList{
		size:   10,
		length: 5,
		data:   [10]int{1, 2, 3, 4, 5},
	}
	s1.Insert(2, 5)
	fmt.Println(s1)

}
