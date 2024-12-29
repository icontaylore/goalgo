package main

import "fmt"

func main() {
	arr := []int{1, 5, 2, 4, 1}
	count := 0

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			fmt.Println(arr[i], i)
			arr[i+1] += (arr[i] - arr[i+1]) + 1
			count++
		}
	}
	fmt.Println(count)
}
