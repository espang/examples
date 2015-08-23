package median

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func _TestRollingMedian(t *testing.T) {
	x := []float64{
		1., 2., 3., 4., 5., 6., 7., 8., 9., 10., 11., 12., 13., 14.,
	}
	h := 2
	rm := RollingMedian(h, x)
	fmt.Println(rm)
}

func _TestRollingMedianWithNaN(t *testing.T) {
	x := []float64{1., 2., 3., math.NaN(), 3., 2.}
	exp := []float64{math.NaN(), 2., 2., 3., 2., math.NaN()}
	h := 1
	rm := RollingMedian(h, x)
	fmt.Println(exp)
	fmt.Println(rm)
}

func benchAlgo(size int, h int, b *testing.B) {
	array := make([]float64, 0, size)
	for i := 0; i < size; i++ {
		array = append(array, rand.Float64())
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		RollingMedian(h, array)
	}
}

func BenchmarkAlgo_100_2(b *testing.B)     { benchAlgo(100, 2, b) }
func BenchmarkAlgo_100_4(b *testing.B)     { benchAlgo(100, 4, b) }
func BenchmarkAlgo_1000_2(b *testing.B)    { benchAlgo(1000, 2, b) }
func BenchmarkAlgo_1000_4(b *testing.B)    { benchAlgo(1000, 4, b) }
func BenchmarkAlgo_10000_2(b *testing.B)   { benchAlgo(10000, 2, b) }
func BenchmarkAlgo_10000_4(b *testing.B)   { benchAlgo(10000, 4, b) }
func BenchmarkAlgo_10000_6(b *testing.B)   { benchAlgo(10000, 6, b) }
func BenchmarkAlgo_10000_8(b *testing.B)   { benchAlgo(10000, 8, b) }
func BenchmarkAlgo_100000_2(b *testing.B)  { benchAlgo(100000, 2, b) }
func BenchmarkAlgo_100000_4(b *testing.B)  { benchAlgo(100000, 4, b) }
func BenchmarkAlgo_100000_6(b *testing.B)  { benchAlgo(100000, 6, b) }
func BenchmarkAlgo_100000_8(b *testing.B)  { benchAlgo(100000, 8, b) }
func BenchmarkAlgo_100000_10(b *testing.B) { benchAlgo(100000, 10, b) }
func BenchmarkAlgo_100000_12(b *testing.B) { benchAlgo(100000, 12, b) }
func BenchmarkAlgo_100000_14(b *testing.B) { benchAlgo(100000, 14, b) }
