package CommonHeapUtilitiesPackage

import "math"

type HeapUtilities struct{}

func (m *HeapUtilities) LeftChild(index int) int {
	return 2*index + 1
}

func (m *HeapUtilities) RightChild(index int) int {
	return 2*index + 2
}

func (m *HeapUtilities) Parent(index int) int {
	return int(math.Floor(float64(index-1) / 2))
}

func (m *HeapUtilities) Swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
