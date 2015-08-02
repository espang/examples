package heap

import (
	"errors"
	"math"
)

type minHeap struct {
	heap
}

func NewMinHeap() *minHeap {
	m := &minHeap{}
	m.vals = append(m.vals, math.NaN())
	m.n = 0
	return m
}

func (m *minHeap) Min() (float64, error) {
	return m.rootValue()
}

func (m *minHeap) DequeMin() (float64, error) {
	if len(m.vals) == 1 {
		return 0, errors.New("Called Pop() on empty heap")
	}
	val := m.vals[1]
	m.swap(1, m.n)
	m.decrease()
	if m.n > 1 {
		m.downheap(1)
	}
	return val, nil
}

func (m *minHeap) Enqueue(v float64) {
	m.vals = append(m.vals, v)
	m.n += 1
	i := m.n
	m.upheap(i)
}

func (m *minHeap) upheap(i int64) {
	if i == 1 {
		return
	}
	j := parent(i)
	if m.vals[i] < m.vals[j] {
		m.swap(i, j)
		m.upheap(j)
	}
}

func (m *minHeap) downheap(i int64) {
	l := left(i)
	if l > m.n {
		return
	}
	r := right(i)
	var j int64
	if r > m.n {
		j = l
	} else if m.vals[l] > m.vals[r] {
		j = r
	} else {
		j = l
	}
	if m.vals[j] < m.vals[i] {
		m.swap(i, j)
		m.downheap(j)
	}
}
