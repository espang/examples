//Give 2 strings:
// Write function to decide wether the two strings are permuations
package datatypes

import "sort"

func IsPermuation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		// Permutations have the same
		// number of letters
		return false
	}

	r1 := RuneSlice([]rune(s1))
	r2 := RuneSlice([]rune(s2))

	sort.Sort(r1)
	sort.Sort(r2)

	return allEqual(r1, r2)
}

func allEqual(slice1, slice2 RuneSlice) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

type RuneSlice []rune

func (p RuneSlice) Len() int {
	return len(p)
}
func (p RuneSlice) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p RuneSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
