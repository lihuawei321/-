package main

import "fmt"

//数组中选择两个数的和等于目标值的下标

func res(arr []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		b := target - arr[i]
		if _, ok := m[b]; ok {
			return []int{m[b], i}
		}
		m[arr[i]] = i
	}
	return nil
}

func main() {
	arr := []int{2, 7, 11, 15}
	ints := res(arr, 9)
	fmt.Println(ints)
}
