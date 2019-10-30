package main

import "fmt"

func gcd(a, b int) int {
	r := a % b
	if r == 0 {
		return b
	} else {
		return gcd(b, r)
	}
}

func main() {
	var a int
	var b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	i := gcd(a, b)
	fmt.Println(i)

}
