package main

import "fmt"

func main() {
	a := []int{23, 3, 1, 2}
	b := []int{6, 2, 4, 23}
	// [2, 23]
	fmt.Printf("%v\n", intersection(a, b))
	a = []int{1, 1, 1}
	b = []int{1, 1, 1, 1}
	// [1, 1, 1]
	fmt.Printf("%v\n", intersection(a, b))
}

func intersection(a, b []int) []int {
	arr, j := []int{}, 0
	sort(a)
	sort(b)

	for i := 0; i < len(a); {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			arr = append(arr, a[i])
			i++
			j++
		}
	}
	return arr
}

func sort(arr []int) {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
