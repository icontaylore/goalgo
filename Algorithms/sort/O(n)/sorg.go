package O_n_

import "fmt"

func main() {
	arr...
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
