package main

import (
	"fmt"
)

// Binary search O(n)

func main() {
	array := []int{3, 4, 9, 15, 19, 22, 27, 54, 76, 100, 104, 133, 144, 155, 187}
	bsearch(array, 100)

}

func bsearch(array []int, i int) {
	left := 0
	right := len(array) - 1

	for j := 0; j < 6; j++ {
		mid := (left + right) / 2
		fmt.Println(mid, j)

		if array[mid] < i {
			fmt.Println(left)
			left = mid + 1
		} else if array[mid] > i {
			right = mid - 1
		}
	}
}
