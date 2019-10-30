package main

import "fmt"

func f(n int) int {
	if n == 1 {
		return 1
	} else {
		return f(n-1) + 2
	}
}

func fib(n int)int {
	if n==1 || n==2{
		return 1
	}else{
		return fib(n-1)+fib(n-2)
	}
	
}

//计算1+2+3+。。。+100
func cal(n int) int {
	if n==1{
		return 1
	}else{
		return cal(n-1)+n
	}
}

//求数组前n元素的和
func sum(arr []int,n int)int  {
	if n==0{
		return arr[0]
	}else{
		return sum(arr,n-1)+arr[n]
	}
	
}
func main() {
	res := f(3)
	fmt.Println(res)
	res2:=fib(6)
	fmt.Println(res2)
	res3:=cal(100)
	fmt.Println(res3)

	arr:=[]int{1,3,5,6,8}
	num1 := sum(arr, 3)
	fmt.Println(num1)
}
