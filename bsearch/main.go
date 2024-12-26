package bsearch

import "fmt"

func bsearch() {
	array := []int{1, 1, 4, 4, 8, 120}
	array2 := []int{1, 2, 4, 5, 6, 7, 8, 63, 64, 65}

	for _, v := range array2 {
		l, r := lbs(array, v), rbs(array, v)
		if array[l]-v < v-array[r] {
			fmt.Println(array[l])
		} else {
			fmt.Println(array[r])
		}
	}
}

func lbs(array []int, i int) int { // r=2;l=-1;m=1/2=0
	l := -1
	r := len(array) - 1

	for l+1 != r {
		mid := (r + l) / 2

		if array[mid] < i {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}

func rbs(array []int, i int) int { //r=2;l=0;m=1;<l++;l=1+1=2exretl
	l := 0
	r := len(array) - 1

	for l+1 != r {
		mid := (r + l) / 2

		if array[mid] < i { //l=7;r=8
			l = mid
		} else {
			r = mid
		}
	}
	return l
}
