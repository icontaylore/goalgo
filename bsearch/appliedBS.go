package bsearch

import "fmt"

func main() {
	findnegative()
}

func findnegative() {
	a, b := 0, 0
	fmt.Scan(&a, &b)
	res := 0

	array := make([]int, b)
	for i := 0; i < a; i++ {

		for j := 0; j < b; j++ {
			fmt.Scan(&array[j])
		}

		if array[len(array)-1] >= 0 {
			res += 0
		} else {
			bsearch(&array, &res)
		}
	}
	fmt.Println(res)
}

func bsearch(array *[]int, res *int) {
	left := 0
	right := len(*array) - 1

	if (*array)[0] >= 0 {
		for left+1 != right {
			mid := (left + right) / 2

			if (*array)[mid] >= 0 {
				left = mid
			} else {
				right = mid
			}
		}

		*res += len(*array) - right
		return
	} else {
		*res += len(*array)
	}
}
