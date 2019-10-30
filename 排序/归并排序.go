package main

import "fmt"

func merge(arr []int, l, m, r int) {
	leftSize := m - l
	rightSize := r - m+1
	left := make([]int, leftSize)
	right := make([]int, rightSize)
	var i,j,k int
	//res:=[]int{}
	for i = l; i < m; i++ {
		left[i-l] = arr[i]
	}
	for i = m; i <= r; i++ {
		right[i-m] = arr[i]
	}
	//
	//fmt.Println(left)
	//fmt.Println(right)
	i = 0
	j = 0
	k = l
	//res:=[]int{}
	for i < leftSize && j < rightSize {
		if left[i] < right[j] {
			arr[k] = left[i]
			//res=append(res,left[i])
			i++
			k++
		}else {
			arr[k] = right[j]
			//res=append(res,right[j])
			j++
			k++
		}

	}
	//fmt.Println(arr)
	for i < leftSize {
		arr[k] = left[i]
		i++
		k++
	}
	for j < rightSize {
		arr[k] = right[j]
		j++
		k++
	}

}
func mergeSort(arr []int, l int, r int) {
	if l<r {
		m := (l + r) / 2
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)
		merge(arr, l, m+1, r)
	}
}

func main() {
	//arr := []int{2, 8, 9, 10, 4, 5, 6, 7}
	arr := []int{6, 8, 10, 9, 4, 5, 2, 7}
	l := 0
	r := len(arr)-1
	//m:=4
	//fmt.Println(r)
	//m := 4
	mergeSort(arr, l, r)
	//merge(arr,l,m,r)
	fmt.Println(arr)
}

