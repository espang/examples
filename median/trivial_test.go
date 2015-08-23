package median

import (
	"fmt"
	"math/rand"
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

func benchTrivial(size int, h int, b *testing.B) {
	array := make([]float64, 0, size)
	for i := 0; i < size; i++ {
		array = append(array, rand.Float64())
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RollingMedianTrivial(h, array)
	}
}

func BenchmarkTrivial_100_2(b *testing.B)     { benchTrivial(100, 2, b) }
func BenchmarkTrivial_100_4(b *testing.B)     { benchTrivial(100, 4, b) }
func BenchmarkTrivial_1000_2(b *testing.B)    { benchTrivial(1000, 2, b) }
func BenchmarkTrivial_1000_4(b *testing.B)    { benchTrivial(1000, 4, b) }
func BenchmarkTrivial_10000_2(b *testing.B)   { benchTrivial(10000, 2, b) }
func BenchmarkTrivial_10000_4(b *testing.B)   { benchTrivial(10000, 4, b) }
func BenchmarkTrivial_10000_6(b *testing.B)   { benchTrivial(10000, 6, b) }
func BenchmarkTrivial_10000_8(b *testing.B)   { benchTrivial(10000, 8, b) }
func BenchmarkTrivial_100000_2(b *testing.B)  { benchTrivial(100000, 2, b) }
func BenchmarkTrivial_100000_4(b *testing.B)  { benchTrivial(100000, 4, b) }
func BenchmarkTrivial_100000_6(b *testing.B)  { benchTrivial(100000, 6, b) }
func BenchmarkTrivial_100000_8(b *testing.B)  { benchTrivial(100000, 8, b) }
func BenchmarkTrivial_100000_10(b *testing.B) { benchTrivial(100000, 10, b) }
func BenchmarkTrivial_100000_12(b *testing.B) { benchTrivial(100000, 12, b) }
func BenchmarkTrivial_100000_14(b *testing.B) { benchTrivial(100000, 14, b) }
