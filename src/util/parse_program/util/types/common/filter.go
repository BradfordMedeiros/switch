package common

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


func FilterArray(arrString []string, filter func(string) bool) []string {
	arr := make([]string, 0)
	for i := 0; i < len(arrString); i++ {
		if filter(arrString[i]){
			arr = append(arr, arrString[i])
		}
	}
	return arr
}