package main

import "fmt"

func hanoi(n int, A, B, C byte) {
	if n == 1 {
		fmt.Printf("%c->%c\n", A, C)
	} else {
		hanoi(n-1, A, C, B)
		fmt.Printf("%c->%c\n", A, C)
		hanoi(n-1, B, A, C)
	}

}

func main() {
	hanoi(3, 'A', 'B', 'C')
}
