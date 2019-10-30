package main

import "fmt"

func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] > insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
	}

}

func InsertSort2(arr []int)  {
	for i:=1;i<len(arr);i++{
		for j:=i;j>0;j--{
			if arr[j]<arr[j-1]{
				arr[j],arr[j-1]=arr[j-1],arr[j]
			}else{
				break
			}
		}
	}

}
func main() {
	arr:=[]int{81,94,11,96,12,35,17,95,28,58,41,75,15}
	InsertSort2(arr)
	fmt.Println(arr)
}
