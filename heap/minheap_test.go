package heap

import "testing"

func (m *minHeap) checkMinHeapDownFrom(i int64) bool {
	l := left(i)
	r := right(i)

	if l > m.n {
		return true
	}
	if r > m.n {
		return m.vals[l] >= m.vals[i]
	}
	if m.vals[l] < m.vals[i] || m.vals[r] < m.vals[i] {
		return false
	}
	return m.checkMinHeapDownFrom(l) && m.checkMinHeapDownFrom(r)
}

func TestIsHeap(t *testing.T) {
	m := NewMinHeap()
	m.Enqueue(4.)
	m.Enqueue(2.)
	m.Enqueue(6.)
	if !m.checkMinHeapDownFrom(1) {
		t.Errorf("No heap")
	}
}

func TestIsHeapFunc(t *testing.T) {
	m := NewMinHeap()
	m.Enqueue(4.)
	m.Enqueue(2.)
	m.Enqueue(6.)
	m.swap(1, 2)
	if m.checkMinHeapDownFrom(1) {
		t.Errorf("No heap because of invalid swap")
	}
	m.swap(1, 2)
	m.swap(1, 3)
	if m.checkMinHeapDownFrom(1) {
		t.Errorf("No heap because of invalid swap")
	}
}
