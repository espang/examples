package heap

type MaxHeap struct {
	heap
}

func (m *MaxHeap) Max() (float64, error) {
	return m.rootValue()
}
