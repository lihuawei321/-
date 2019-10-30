package main

import (
	"errors"
)

type SequenStack struct {
	Data []interface{}
}

func New() *SequenStack {
	s := make([]interface{}, 0)
	return &SequenStack{
		s,
	}

}

//压栈
func (ss *SequenStack) Push(data interface{}) {
	ss.Data = append(ss.Data, data)

}

//出栈
func (ss *SequenStack) Pop() (interface{}, error) {
	if ss.Length() == 0 {
		return nil, errors.New("stack is empty")
	}
	index := ss.Length() - 1
	value := ss.Data[index]
	ss.Data = append(ss.Data[:index])
	return value, nil
}

//获取栈的长度
func (ss *SequenStack) Length() int {
	return len(ss.Data)

}
func main() {

}
