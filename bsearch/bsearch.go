package bsearch

import "fmt"

// Binary search O(n)

func main() {
	array := []int{3, 4, 6, 8, 13, 66, 88, 99, 102, 144, 166}

	i := 144
	l := 0
	r := len(array) - 1

	for l <= r {
		mid := (l + r) / 2

		if array[mid] == i {
			fmt.Println(mid)
			return
		} else if array[mid] < i {
			l = mid
		} else {
			r = mid
		}
	}
}
