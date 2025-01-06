package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	if n == 0 {
		fmt.Println(dp[0])
	} else {
		for i := 2; i < len(dp); i++ { // 0 (0); 1(1); 2(1); 3(2); 4(3);
			dp[i] = dp[i-1] + dp[i-2]
			fmt.Println(i, "i", dp[i])
		}
	}
	fmt.Println(dp[n])
}
