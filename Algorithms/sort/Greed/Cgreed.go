package main

import (
	"fmt"
	"sort"
)

// 6 615
// 140 215 120 160 100 340

func main() {
	w := 615
	arr := []int{140, 215, 120, 160, 100, 340} // 6 грузов
	count := 0
	col := 0
	newSortArr := []int{}

	for i, _ := range arr {
		if arr[i] >= 210 && arr[i] <= 220 {
			count += arr[i]
			col++
		} else {
			newSortArr = append(newSortArr, arr[i])
		}
	}
	// sort
	sort.Sort(sort.Reverse(sort.IntSlice(newSortArr)))
	//
	for i := 0; i < len(newSortArr); i++ {
		if count+newSortArr[i] <= w {
			fmt.Println(count)
			for j := 1; j < len(newSortArr); j++ {
				if count+newSortArr[j] <= w {
					count += newSortArr[j]
					col++
				}
			}
		} else {
			for l := 1; l < len(newSortArr); l++ {
				if count+newSortArr[l] <= w {
					count += newSortArr[l]
					col++
				}
			}
		}
	}
	fmt.Println(col, count)
}
