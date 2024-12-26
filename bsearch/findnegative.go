package bsearch

import "fmt"

func main() {
	array := []int{1, 1, 1, -2}

	res := 0
	if array[0] >= 0 {
		res = len(array) - bs(array)

	} else {
		res += len(array)
	}
	fmt.Println(res)
}

func bs(arr []int) int {
	//arr := []int{} //int
	left := 0             //0
	right := len(arr) - 1 //7/2=3

	for left+1 != right {
		mid := (left + right) / 2 // center = 5/2 = 2

		if arr[mid] >= 0 {
			left = mid // left = 2(3); right = 5
		} else { // l=1(2);r=3;4/2=2; r=2;r=1(2) / end cicle
			right = mid
		}
	}
	return right
}
