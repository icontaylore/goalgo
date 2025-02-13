package main

import "fmt"

func main() {
	data := []int{-5, 10, 8, 10, 2, -3, 10}

	arr := make([]int, len(data), len(data)*2)
	for i := 0; i < len(data); i++ {
		arr[i] = data[i]
	}

	fmt.Println(-8 > -16)
	flag := 0
	for i := 0; i < len(arr); i++ {
		if flag == len(arr) {
			break
		}
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			i = -1
			flag = 0
		}
		flag++
	}
	fmt.Println(arr)
}

//
//
//func main() {
//	data := []int{-8, 42, 18, 0, -16}
//
//	arr := make([]int, len(data), len(data)*2)
//	for {
//		minus := 1
//		for i := 0; i < len(data); i += 2 {
//			if i+1 == len(data) {
//				arr[len(arr)-minus] = data[i]
//				break
//			}
//			if data[i] < data[i+1] {
//				arr[i] = data[i]
//				arr[len(arr)-minus] = data[i+1]
//				minus++
//			} else if data[i] > data[i+1] {
//				arr[len(arr)-minus] = data[i]
//				arr[i] = data[i+1]
//				minus++
//			} else {
//				break
//			}
//		}
//	}
//	fmt.Println(arr)
//}
