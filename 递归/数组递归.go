package main

import "fmt"

//找数组的最大值
func findMax(arr []int,l,r int)  int{
	if l==r{
		return arr[l]
		}else{
			a:=arr[l]
			b:=findMax(arr,l+1,r)
			if a>b{
				return a
			}else{
				return b
			}
	}
}

//数组求和
func arrSum(arr []int,l,r int)int{
	if l==r{
		return arr[l]
	}else {
		return arr[l]+arrSum(arr,l+1,r)
	}
}

//冒泡排序
func bubbleSort(arr []int,l,r int)  {
	if l<r{
		for i:=l;i<=r-1;i++{
			if arr[i]>arr[i+1]{
				arr[i],arr[i+1]=arr[i+1],arr[i]
			}
		}
		bubbleSort(arr,l,r-1)
	}
}

func main() {
	arr:=[]int{2,7,3,1,6,5}
	max := findMax(arr, 0, 5)
	fmt.Println(max)
	i := arrSum(arr, 0, 5)
	fmt.Println(i)
	bubbleSort(arr,0,5)
	fmt.Println(arr)
}
