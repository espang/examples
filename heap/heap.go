package heap

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// heap as an array
//    1
//  3    5
// 4 5  6 7
// -> x 1 3 5 4 5 6 7
//    0 1 2 3 4 5 6 7
// Min/ Max in vals[1]
// children: 2*i 2*i+1
// parent: i/2
type heap struct {
	vals []float64
	n    int64
}

func left(i int64) int64   { return 2 * i }
func right(i int64) int64  { return 2*i + 1 }
func parent(i int64) int64 { return i / 2 }

// swap elements from indices i and j
// no error checking is intended
// because the method is not exported
func (h *heap) swap(i, j int64) {
	h.vals[i], h.vals[j] = h.vals[j], h.vals[i]
}

func (h *heap) decrease() {
	if h.n > 1 {
		h.vals = h.vals[:h.n]
		h.n -= 1
	}
}

// Pop removes the root (Minimum) of
// the Minheap and returns it
func (h *heap) pop() (float64, error) {
	if len(h.vals) == 1 {
		return 0, errors.New("Called Pop() on empty heap")
	}
	val := h.vals[1]
	h.swap(1, h.n)
	h.decrease()
	if h.n > 1 {
		h.downheap(1)
	}
	return val, nil
}

func (h *heap) rootValue() (float64, error) {
	if len(h.vals) > 1 {
		return h.vals[1], nil
	}
	return 0, errors.New("Failed to get the root value on empty heap")
}

func (h *heap) downheap(i int64) {}

func (h *heap) Print() {
	fmt.Println("Heap with ", h.n, "elements")
	fmt.Printf("%v\n", h)
	levels := int(math.Ceil(math.Log(float64(h.n)))) + 1
	for i := 1; i <= levels; i++ {
		h.printline(i, levels)
	}
}

func pow(base, exponent int) int {
	if exponent == 0 {
		return 1
	} else if exponent < 0 {
		panic("Exponent is smaller than zero, not implemented")
	}
	return base * pow(base, exponent-1)
}

func tostring(vs []float64) []string {
	res := make([]string, 0, len(vs))
	for _, v := range vs {
		res = append(res, fmt.Sprintf("%.2f", v))
	}
	return res
}

// line 1: [1:2] <=> [2^0 : 2^1]
// line 2: [2:4]
// line 3: [4:8]
// line 4: [8:16]
func (h *heap) printline(i, maxlevel int) {
	lineLength := pow(2, maxlevel)
	start := pow(2, i-1)
	end := int(math.Min(float64(h.n)+1, float64(pow(2, i))))
	ss := tostring(h.vals[start:end])
	shift := strings.Repeat(" ", 2+lineLength/pow(2, i))
	fmt.Println("level: ", i, " ::", shift, strings.Join(ss, "  "))
}
