package main

func Decode(roman string) int {
	str := roman
	mapa := make(map[string]int)
	mapa["I"], mapa["V"], mapa["X"], mapa["L"], mapa["C"], mapa["D"], mapa["M"] = 1, 5, 10, 50, 100, 500, 1000

	res := 0

	old := mapa[string(str[0])]
	now := 0

	for i := 1; i < len(str); i++ {
		now = mapa[string(str[i])]
		if old < now {
			res -= old
		} else if old >= now {
			res += old
		}
		old = now
	}

	// default
	res += mapa[string(str[len(str)-1])]
	return res
}
