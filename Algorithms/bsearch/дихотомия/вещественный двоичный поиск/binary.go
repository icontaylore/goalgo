package вещественный_двоичный_поиск

import (
	"fmt"
	"math"
)

func main() {
	a, n := float64(2), 2

	var l, r float64 = 0, 1000
	e := 1e-7 // эпсилон (0.0000000..)
	for float64(r)-float64(l) > e {
		mid := float64((l + r) / 2)

		if math.Pow(mid, float64(n)) < a {
			l = mid
		} else {
			r = mid
		}
	}
	fmt.Println(l)
}
