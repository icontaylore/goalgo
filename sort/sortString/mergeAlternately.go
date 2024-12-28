package sortString

func main() {
	mergeAlternately("ab", "pqrs")

}

func mergeAlternately(word1 string, word2 string) string {
	res := ""
	i, j := 0, 0

	for i < len(word1) || j < len(word2) {
		if i < len(word1) {
			res += string(word1[i])
		}

		if j < len(word2) {
			res += string(word2[i])
		}
		i++
		j++
	}
	return res
}
