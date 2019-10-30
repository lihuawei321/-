package main

import "fmt"

func ShellSort(arr []int) {
	len:=len(arr)
	for d := len / 2; d >0; d /= 2 {
		for i:=d;i<len;i++{
			for j:=i;j>=d;j-=d{
				if arr[j]<arr[j-d] {
					arr[j], arr[j-d] = arr[j-d], arr[j]
				}else{
					break
				}
			}
		}
	}
}

func main() {
	arr:=[]int{81,94,11,96,12,35,17,95,28,58,41,75,15}
	ShellSort(arr)
	fmt.Println(arr)
}
