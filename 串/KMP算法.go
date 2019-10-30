package main

import "fmt"

func kmp(haystack string, needle string) int {
	//when needle is empty,return 0 is available
	if len(needle) == 0 {
		return 0
	}

	next := getNextBetter(needle)

	i := 0
	j := 0
	for i < len(haystack) && j < len(needle) {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == len(needle) {
		return i - j
	}

	return -1
}

//original getNext function
func getNextBetter(str string) []int {
	var next = make([]int, len(str))
	next[0] = -1

	i := 0
	j := -1
	for i < len(str)-1 {
		if j == -1 || str[i] == str[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}

	}

	return next
}

//better getNext function
func getNext(str string) []int {
	var next = make([]int, len(str))
	next[0] = -1

	i := 0
	j := -1
	for i < len(str)-1 {
		if j == -1 || str[i] == str[j] {
			i++
			j++
			if str[i] == str[j] {
				next[i] = next[j]
			} else {
				next[i] = j
			}
		} else {
			j = next[j]
		}
	}

	return next
}

func main() {
	a := "abcababaca"
	b := "abac"
	d := getNextBetter(b)
	c := kmp(a, b)
	fmt.Println(d)
	fmt.Println(c)
}
