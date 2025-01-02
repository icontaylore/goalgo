package main

import "fmt"

func main() {
	n := 5
	arr := []int{0, 1, 2, 1, 4}

	res := 0
	end := 0
	// n*n=n^2
	for i := 0; i < n; i++ { // 0+1
		res -= arr[i]
		for j := 0; j < n; j++ {
			res += arr[j]
			if res == arr[i] {
				end++
			}
		}
		res = 0
	}
	fmt.Println(end)
}
