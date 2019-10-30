package main

import "fmt"

func heapSort(arr []int)  {
	arrLen := len(arr)
	buildMaxHeap(arr, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		swap(arr, 0, i)
		//arrLen -= 1
		heapify(arr, 0, i)
	}
	return
}

func buildMaxHeap(arr []int, arrLen int) {
	for i := (arrLen - 1) / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}

}

func heapify(arr []int, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, arrLen)
	}

}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]

}

func main() {
	arr := []int{5, 2, 7, 3, 6, 1, 4}
	//arr:=[]int{4,10,3,5,1,2}
	//heapify(arr,0,7)
	heapSort(arr)
	fmt.Println(arr)
}
