package main

func QuickSort(left, right int, arr []int) {
	l := left
	r := right
	pivot := arr[(left+right)/2]
	for l < r {
		for arr[l] < pivot {
			l++
		}
		for arr[r] > pivot {
			r--
		}
		if l > r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		QuickSort(left, r, arr)
	}
	if right > l {
		QuickSort(l, right, arr)
	}

}

func main() {

}
