package datatypes

import "testing"

type UniqueFunc func(string) bool

var funcmap = map[string]UniqueFunc{
	"IsUniqueTrivial": IsUniqueTrivial,
	"IsUnique":        IsUnique,
	"IsUnique2":       IsUnique2,
}

var stringmap = map[string]bool{
	"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890":  true,
	"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890a": false,
	"好\n\t\\":     true,
	"好\n\t\n\\好":  false,
	"好\n\t\n\\好a": false,
	"PÐŐǐ":        true,
}

func TestIsUnique(t *testing.T) {
	for s, result := range stringmap {
		for name, f := range funcmap {
			if f(s) != result {
				t.Errorf("%s returns false for a unique string: '%s'", name, s)
			}
		}
	}
}

func TestIsPermutation(t *testing.T) {
	s1 := "asdf"
	s2 := "sdaf"
	if !IsPermuation(s1, s2) {
		t.Errorf("IsPermuation returns false for permutations")
	}
	s3 := "asdfa"
	if IsPermuation(s1, s3) {
		t.Errorf("IsPermuation returns true but should not: ('%s' - '%s')", s1, s3)
	}
	s4 := "asdd"
	if IsPermuation(s1, s4) {
		t.Errorf("IsPermuation returns true but should not: ('%s' - '%s')", s1, s4)
	}
}

func TestUrlify(t *testing.T) {
	s := "Mr John Smith    "
	url := "Mr%20John%20Smith"
	res := Urlify(s, 13)
	if res != url {
		t.Errorf("Urlify(%s) -> %s, should be %s", s, res, url)
	}
	s = " Mr John Smith      "
	url = "%20Mr%20John%20Smith"
	res = Urlify(s, 14)
	if res != url {
		t.Errorf("Urlify(%s) -> %s, should be %s", s, res, url)
	}
}

func TestPalindrom(t *testing.T) {
	s1 := "taco cat"
	if !IsPermutationOfPalindroms(s1) {
		t.Errorf("IsPermutationOfPalindroms false, but should be true")
	}
	s1 = "taco catt"
	if IsPermutationOfPalindroms(s1) {
		t.Errorf("IsPermutationOfPalindroms true, but should be false")
	}
}

func TestRotate2x2(t *testing.T) {
	p1 := NewPixel('a')
	p2 := NewPixel('b')
	p3 := NewPixel('c')
	p4 := NewPixel('d')
	image := &Image{data: make([][]Pixel, 0, 2), n: 2}
	image.data = append(image.data, []Pixel{p1, p2}, []Pixel{p3, p4})
	rotate := &Image{data: make([][]Pixel, 0, 2), n: 2}
	rotate.data = append(rotate.data, []Pixel{p2, p4}, []Pixel{p1, p3})
	RotateMatrix(image)
	if !image.Equals(rotate) {
		t.Errorf("RotateMatrix not right")
	}
}

func TestRotate3x3(t *testing.T) {
	p1 := NewPixel('a')
	p2 := NewPixel('b')
	p3 := NewPixel('c')
	p4 := NewPixel('d')
	p5 := NewPixel('e')
	p6 := NewPixel('f')
	p7 := NewPixel('g')
	p8 := NewPixel('h')
	p9 := NewPixel('i')
	image := &Image{data: make([][]Pixel, 0, 3), n: 3}
	image.data = append(
		image.data,
		[]Pixel{p1, p2, p3},
		[]Pixel{p4, p5, p6},
		[]Pixel{p7, p8, p9},
	)
	rotate := &Image{data: make([][]Pixel, 0, 3), n: 3}
	rotate.data = append(
		rotate.data,
		[]Pixel{p3, p6, p9},
		[]Pixel{p2, p5, p8},
		[]Pixel{p1, p4, p7},
	)
	RotateMatrix(image)
	if !image.Equals(rotate) {
		t.Errorf("RotateMatrix not right")
	}
}

func TestRotate4x4(t *testing.T) {
	p1 := NewPixel('a')
	p2 := NewPixel('b')
	p3 := NewPixel('c')
	p4 := NewPixel('d')
	p5 := NewPixel('e')
	p6 := NewPixel('f')
	p7 := NewPixel('g')
	p8 := NewPixel('h')
	p9 := NewPixel('i')
	p10 := NewPixel('j')
	p11 := NewPixel('k')
	p12 := NewPixel('l')
	p13 := NewPixel('m')
	p14 := NewPixel('n')
	p15 := NewPixel('o')
	p16 := NewPixel('p')
	image := &Image{data: make([][]Pixel, 0, 4), n: 4}
	image.data = append(
		image.data,
		[]Pixel{p1, p2, p3, p4},
		[]Pixel{p5, p6, p7, p8},
		[]Pixel{p9, p10, p11, p12},
		[]Pixel{p13, p14, p15, p16},
	)
	rotate := &Image{data: make([][]Pixel, 0, 4), n: 4}
	rotate.data = append(
		rotate.data,
		[]Pixel{p4, p8, p12, p16},
		[]Pixel{p3, p7, p11, p15},
		[]Pixel{p2, p6, p10, p14},
		[]Pixel{p1, p5, p9, p13},
	)
	RotateMatrix(image)
	if !image.Equals(rotate) {
		t.Errorf("RotateMatrix not right")
	}

}
