// Code from "Median Filtering is Equivalent to Sorting"
// by Jukka Suomela
// see http://arxiv.org/pdf/1406.1717v1.pdf
package median

import "math"

type idxVal struct {
	v float64
	i int
}

func (iv idxVal) less(cmp idxVal) bool {
	return iv.v < cmp.v
}

type idxVals []idxVal

func (vs idxVals) Len() int {
	return len(vs)
}

func (vs idxVals) Swap(i, j int) {
	vs[i], vs[j] = vs[j], vs[i]
}

func (vs idxVals) Less(i, j int) bool {
	return vs[i].less(vs[j])
}

type block struct {
	k     int
	alpha []float64
	pi    []int
	prev  []int
	next  []int
	tail  int
	m     int
	s     int
}

func newBlock(h int, alpha []float64) *block {
	k := len(alpha)
	pi := sort_block(alpha)
	prev := make([]int, k+1)
	next := make([]int, k+1)
	tail := k
	m := pi[h]
	s := h
	b := &block{k, alpha, pi, prev, next, tail, m, s}
	b.initLinks()
	return b
}

func (b *block) initLinks() {
	p := b.tail
	var q int
	for i := 0; i < b.k; i++ {
		q = b.pi[i]
		b.next[p] = q
		b.prev[q] = p
		p = q
	}
	b.next[p] = b.tail
	b.prev[b.tail] = p
}
func (b *block) unwind() {
	for i := b.k - 1; i > -1; i-- {
		b.next[b.prev[i]] = b.next[i]
		b.prev[b.next[i]] = b.prev[i]
	}
	b.m = b.tail
	b.s = 0
}

func (b *block) atEnd() bool {
	return b.m == b.tail
}

func (b *block) getPair(i int) idxVal {
	return idxVal{b.alpha[i], i}
}
func (b *block) isSmall(i int) bool {
	return b.atEnd() || b.getPair(i).less(b.getPair(b.m))
}

func (b *block) delete(i int) {
	b.next[b.prev[i]] = b.next[i]
	b.prev[b.next[i]] = b.prev[i]
	if b.isSmall(i) {
		b.s -= 1
	} else {
		if b.m == i {
			b.m = b.next[b.m]
		}
		if b.s > 0 {
			b.m = b.prev[b.m]
			b.s -= 1
		}
	}
}

func (b *block) undelete(i int) {
	b.next[b.prev[i]] = i
	b.prev[b.next[i]] = i
	if b.isSmall(i) {
		b.m = b.prev[b.m]
	}
}

func (b *block) advance() {
	b.m = b.next[b.m]
	b.s += 1
}

func (b *block) peek() float64 {
	if b.atEnd() {
		return math.Inf(1)
	}
	return b.alpha[b.m]
}
