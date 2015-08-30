package datatypes

import "unicode"

func Urlify(s string, last int) string {
	url := []rune(s)

	currentIndex := len(s) - 1
	var r rune

	for i := last - 1; i >= 0; i-- {
		r = url[i]
		if !unicode.IsSpace(r) {
			url[currentIndex] = r
			currentIndex--
			continue
		}
		url[currentIndex] = '0'
		url[currentIndex-1] = '2'
		url[currentIndex-2] = '%'
		currentIndex = currentIndex - 3
	}
	return string(url)
}
