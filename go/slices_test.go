package main

import (
	"sort"
	"testing"
)

func TestCopySlices(t *testing.T) {
	s1 := []float64{1., 3., 2.}
	s2 := make([]float64, len(s1), len(s1))
	copy(s2, s1)
	sort.Float64s(s2)
	if !(s1[0] == 1. && s1[1] == 3. && s1[2] == 2.) {
		t.Errorf("Array s1 changed values or order from [1., 3., 2.] to %#v", s1)
	}
	if !(s2[0] == 1. && s2[1] == 2. && s2[2] == 3.) {
		t.Errorf("Array s2 is not sorted: %#v", s2)
	}
}
