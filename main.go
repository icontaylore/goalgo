package main

import (
	"fmt"
)

func main() {
	res := 1 //1+1*2+1
	count := 1

	dp := []int{res}

	if res == 0 { //
		fmt.Println(0)
	} else if res == 2 {
		fmt.Println(1)
	} else {
		for dp[0] > 2 {
			fmt.Println(dp[0])
			if dp[0]%3 == 0 {
				dp[0] /= 3
			} else if dp[0]%2 == 0 {
				dp[0] /= 2
			} else {
				dp[0] -= 1
			}

			count++
		}
		fmt.Println(dp[0], count)
	}
}
