package filter

// typical filter function found in many languages

func Filter(filterString string, filter func(rune) bool) string {
	runes := make([]rune, len(filterString))
	for _, rune := range(filterString){
		if filter(rune) {
			runes = append(runes, rune)
		}
	}
	return string(runes)
}
