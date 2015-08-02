package median

import (
	"log"
	"math"
	"sort"
)

func sort_block(alpha []float64) []int {
	res := idxVals{}
	for i, a := range alpha {
		res = append(res, idxVal{a, i})
	}
	sort.Sort(res)
	si := make([]int, len(alpha))
	for i, iv := range res {
		si[i] = iv.i
	}
	return si
}

func blockChan(h, k, b int, x []float64) <-chan *block {
	c := make(chan *block, b)
	go func() {
		for j := 0; j < b; j++ {
			c <- newBlock(h, x[j*k:(j+1)*k])
		}
		close(c)
	}()
	return c
}

func RollingMedian(h int, x []float64) []float64 {
	k := 2*h + 1
	to_append := len(x) % k
	for i := 0; i < to_append; i++ {
		x = append(x, math.Inf(1))
	}

	b := len(x) / k

	log.Printf("rolling median: h(%d), k(%d), b(%d)", h, k, b)
	blocks := blockChan(h, k, b, x)
	B := <-blocks
	y := make([]float64, 0, k)
	for i := 0; i < h; i++ {
		y = append(y, math.NaN())
	}
	y = append(y, B.peek())

	var A *block
	for b := range blocks {
		A = B
		B = b
		B.unwind()
		for i := 0; i < k; i++ {
			A.delete(i)
			B.undelete(i)
			if A.s+B.s < h {
				if A.peek() <= B.peek() {
					A.advance()
				} else {
					B.advance()
				}
			}
			v := math.Min(A.peek(), B.peek())
			if math.IsInf(v, 1) || math.IsInf(v, -1) {
				v = math.NaN()
			}
			y = append(y, v)
		}
	}
	for i := 0; i < h; i++ {
		y = append(y, math.NaN())
	}

	return y[:len(x)-to_append]
}
