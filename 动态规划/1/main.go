package main

import (
	"fmt"
)

//选不相邻的数字，加起来和最大
//递归，产生重叠子问题
func rec_opt(arr []int, i int) int {
	if i == 0 {
		return arr[0]
	} else if i == 1 {
		return max(arr[0], arr[1])
	} else {
		A := rec_opt(arr, i-2) + arr[i]
		B := rec_opt(arr, i-1)
		return max(A, B)
	}
}

//非递归
func dp_opt(arr []int) int {
	var opt []int
	opt = make([]int, len(arr))
	opt[0] = arr[0]
	opt[1] = max(arr[0], arr[1])
	for i := 2; i < len(arr); i++ {
		A := opt[i-2] + arr[i]
		B := opt[i-1]
		opt[i] = max(A, B)
	}
	return opt[len(arr)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}
func main() {
	arr := []int{1, 2, 4, 1, 7, 8, 3}
	//i := len(arr) - 1
	//opt := rec_opt(arr, i)
	opt := dp_opt(arr)
	fmt.Println(opt)
}
