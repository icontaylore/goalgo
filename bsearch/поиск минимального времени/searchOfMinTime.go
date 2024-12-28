package поиск_минимального_времени

import (
	"fmt"
	"math"
)

func main() {
	n := 5 // copy
	x := 1 // sec of first
	y := 2 // sec of second

	maxInt := math.Min(float64(x), float64(y))
	m := int(maxInt)

	l := 0
	r := (n * m)

	for l+1 != r { // 5+6=11/2=5
		mid := (l + r) / 2 //4 ;l = 0 ; r = 9+4=13/2=6 r = 6;l=(5)

		if (mid/x)+(mid/y) <= n {
			l = mid
		} else {
			r = mid
		}
	}
	fmt.Println(r)
}
