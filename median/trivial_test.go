package median

import (
	"fmt"
	"testing"
)

func TestMedian(t *testing.T) {
	x := []float64{1., 2., 3.}
	is := median(x)
	exp := 2.
	if exp != is {
		t.Errorf("Median of %s is not %f but %f", x, is, exp)
	}
	x = append(x, 4.)
	is = median(x)
	exp = 2.5
	if exp != is {
		t.Errorf("Median of %s is not %f but %f", x, is, exp)
	}
}

func TestTrivial(t *testing.T) {
	x := []float64{
		1., 2., 3., 4., 5., 6., 7., 8., 9., 10., 11., 12., 13., 14.,
	}
	h := 2
	rm := RollingMedianTrivial(h, x)
	fmt.Println(rm)
}

func BenchTrivial(b *testing.B) {

}
