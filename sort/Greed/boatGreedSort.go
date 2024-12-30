package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	str := strings.Split(input, " ")

	arr, arr2 := []int{}, []int{}

	maxPeop := 0
	fmt.Scan(&maxPeop)

	for _, v := range str {
		toVal, err := strconv.Atoi(string(v))
		if err != nil {
			log.Fatal()
		}
		if toVal <= maxPeop {
			arr = append(arr, toVal)
		} else {
			arr2 = append(arr2, toVal)
		}
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	count := 0
	// cylces for sum
	for i := 0; i < len(arr)-1; i += 2 {
		if arr[i]+arr[i+1] > maxPeop {
			count += 2
		} else {
			count += 1
		}
	}
	fmt.Println(arr, count)
}
