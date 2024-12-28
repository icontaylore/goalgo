package bsearch

import "fmt"

func main() {
	arr := []int{5}

	search(arr, -5)
}

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l + r) / 2

		if nums[mid] == target {
			fmt.Println(nums[mid])
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}
	fmt.Println(-1)
	return -1
}
