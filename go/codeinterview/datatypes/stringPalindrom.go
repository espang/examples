package datatypes

import "unicode"

func IsPermutationOfPalindroms(s string) bool {
	letters := map[rune]int{}
	countOdd := 0
	for _, c := range s {
		if unicode.IsSpace(c) {
			continue
		}

		letters[c] += 1
		if letters[c]%2 == 1 {
			countOdd++
		} else {
			countOdd--
		}
	}
	return countOdd <= 1
}
