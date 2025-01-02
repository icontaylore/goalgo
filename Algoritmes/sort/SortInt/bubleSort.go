package main

import "fmt"

func main() {
	bubleSort([]int{55, 3, 1, 7, 18, 44, 28, 21, 22, 100, 23, 16, 17, 2, 39, 61, 62, 60})
}

func bubleSort(arr []int) []int {
	schet := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				fmt.Println(arr)
				schet++
			}
		}
	}
	fmt.Println(schet)
	return arr
}
