// Given a string:
// Write function that returns wether all letters ocure at most one time
package datatypes

// implement isUnique
func IsUniqueTrivial(s string) bool {
	for i, c1 := range s {
		for j, c2 := range s {
			if c1 == c2 && i != j {
				return false
			}
		}
	}
	return true
}

func IsUnique(s string) bool {
	usedChars := make(map[rune]bool)
	for _, char := range s {
		if usedChars[char] {
			return false
		}
		usedChars[char] = true
	}
	return true
}

type element struct {
	val  rune
	next *element
}

type runemap struct {
	elements [128]*element
	length   int32
}

func (rm *runemap) addRuneIfNotExist(r rune) bool {
	index := r % rm.length
	iter := rm.elements[index]
	if iter == nil {
		rm.elements[index] = &element{val: r}
		return true
	}

	var last *element

	for iter != nil {
		if iter.val == r {
			return false
		}
		last, iter = iter, iter.next
	}
	last.next = &element{val: r}
	return true

}

func IsUnique2(s string) bool {
	rm := runemap{length: 128}
	var added bool
	for _, r := range s {
		added = rm.addRuneIfNotExist(r)
		if !added {
			return false
		}
	}
	return true
}
