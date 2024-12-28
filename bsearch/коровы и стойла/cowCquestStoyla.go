package main

import "fmt"

func main() {
	c := 4
	arr := []int{1, 2, 3, 4, 5, 6, 7} //

	search(arr, c)
}

func search(arr []int, i int) { //11*2
	r := len(arr) - 1
	lowest := 0
	outer := arr[r] - arr[0]

	for lowest+1 != r { //
		mid := (lowest + r) / 2 //

		if check(arr[mid], i, outer, arr[0]) == false {
			lowest = mid
		} else {
			r = mid
		}
		fmt.Println(r, "ответ *можно и mid указать*")
	}
}

func check(i int, j int, max int, a0 int) bool {
	res := true

	if (i-a0)*(j-1) < max {
		res = false
	}

	return res
}
