package median

import (
	"math"
	"sort"
)

func median(s []float64) float64 {
	sort.Float64s(s)
	n := len(s)
	if n%2 == 0 {
		return (s[n/2-1] + s[n/2]) / 2.
	}
	return s[n/2]
}

func RollingMedianTrivial(h int, x []float64) []float64 {
	n := len(x)
	res := make([]float64, 0, n)
	for i := 0; i < h; i++ {
		res = append(res, math.NaN())
	}
	for i := h; i < n-h; i++ {
		res = append(res, median(x[i-h:i+h+1]))
	}
	for i := 0; i < h; i++ {
		res = append(res, math.NaN())
	}
	return res
}
