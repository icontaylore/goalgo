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

	count, j := 0, 0
	// cylces for sum
	for i := len(arr) - 1; i >= j; i-- { // i >= j обозначает, что при пересечении j, работа остановится все проверены
		fmt.Println(arr[i], arr[j])
		if arr[i]+arr[j] <= maxPeop {
			count += 1
			j++
		} else {
			count += 1 // 2//j1; 2+2
		}
	}
	fmt.Println(arr, count)
}
