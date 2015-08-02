package median

import (
	"fmt"
	"math"
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
