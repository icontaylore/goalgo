package sortString

func main() {
	// Вызов функции для слияния строк
	mergeAlternately("ab", "pqrs")
}

func mergeAlternately(word1 string, word2 string) string {
	// Инициализация пустой строки для хранения результата
	res := ""
	// Индексы для обоих слов
	i, j := 0, 0

	// Перебор символов в обоих словах
	for i < len(word1) || j < len(word2) {
		// Добавление символа из word1, если есть
		if i < len(word1) {
			res += string(word1[i])
		}

		// Добавление символа из word2, если есть
		if j < len(word2) {
			res += string(word2[j])
		}

		// Увеличение индексов
		i++
		j++
	}
	return res
}
