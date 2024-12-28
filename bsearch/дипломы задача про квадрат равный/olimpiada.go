package дипломы_задача_про_квадрат_равный

import (
	"fmt"
	"math"
)

func main() {
	n := 12 // diploms
	w := 13 // width
	h := 12 // high

	inf := math.Max(float64(w), float64(h)) * float64(n)
	l := 0
	r := int(inf)

	for l+1 != r {
		mid := (l + r) / 2

		if (mid/w)*(mid/h) > n {
			r = mid
		} else {
			l = mid
		}
	}
	fmt.Println(r)
}
