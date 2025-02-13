package main

func main() {
	a := []int{5, 1, 8, 10, 1, 8, -2, 5, 9, 3, 1, 8, 11, 48} // 3 2 1
	b := []int{6, 1, -2, 6, 2, 1, 7}

	mapka := make(map[int]bool)
	for _, num := range b {
		mapka[num] = true
	}

	arr := []int{}
	for _, v := range a {
		if !mapka[v] {
			arr = append(arr, v)
		}
	}
	return arr
}
