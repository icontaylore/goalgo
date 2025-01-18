package вхождение_одного_в_другое

func CheckPlenty(arr1 []int, arr2 []int) []int {
	arr := []int{}
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr2[j] < arr1[i] {
			j++
		} else if arr2[j] > arr1[i] {
			i++
		} else {
			arr = append(arr, arr2[j])
			j++
			i++
		}
	}
	return arr
}
