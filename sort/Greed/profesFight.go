package main

import (
	"fmt"
	"sort"
)

// 3
// 1 5 ; 2 3 ; 3 4

func main() {
	n := 2

	arr2 := [][]int{}
	for i := 0; i < n; i++ {
		arr := make([]int, 2)
		fmt.Scan(&arr[0], &arr[1])
		arr2 = append(arr2, arr)

	}
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i][1] < arr2[j][1]
	})

	j, count := 0, 1
	for i := 0; i < len(arr2)-1; i++ {
		if arr2[i][j+1] <= arr2[i+1][j] {
			count++
		}
	}
	fmt.Println(arr2, count)
}
