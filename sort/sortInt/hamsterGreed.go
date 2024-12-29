package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{490, 125, 320, 144, 215, 50}
	count, bigNum := 0, 0

	newArr := []int{}

	// отбор меньших до 50
	for _, v := range arr {
		if v > 50 {
			newArr = append(newArr, v)
		} else { // все равно добавляем в общую сумму без скидки
			count += v
		}
	}

	// сортируем массив
	sort.Sort(sort.IntSlice(newArr))

	// стакаем сумму чисел которые не входят в скидку
	for i := len(newArr) / 2; i < len(newArr); i++ {
		count += newArr[i]
	}

	// прибавляем товары со скидкой
	for i := 0; i < len(newArr)/2; i++ {
		count += newArr[i] * 75 / 100
		if newArr[i] > bigNum {
			bigNum = newArr[i]
		}
	}
	fmt.Println(count, bigNum)
}
