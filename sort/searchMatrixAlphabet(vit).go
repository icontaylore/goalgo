package main

import "fmt"

func main() {
	a := 0
	fmt.Scan(&a)

	arr := make([]string, 3)
	for i := 0; i < a; i++ {
		fmt.Scan(&arr[i])
	}

	arr2 := []rune{}
	for _, v := range arr {
		for _, v2 := range v {
			arr2 = append(arr2, v2)
		}
	}

	for i := 0; i < a; i++ {
		if arr2[i] > arr2[i+a] {
			i++
		}
	}

}
