package heap

// MaxHeap data structure
type MaxHeap []int

// Insert data to max heap
func (heap *MaxHeap) Insert(data int) {
	if heap.IsEmpty() {
		*heap = append(*heap, data)
		return
	}

	*heap = append(*heap, data)
	index := len(*heap) - 1

	for index != 0 && (*heap)[index] > (*heap)[parent(index)] {
		(*heap)[parent(index)], (*heap)[index] = (*heap)[index], (*heap)[parent(index)]
		index = parent(index)
	}
}

// Heapify updates by requirement of max heap
func (heap *MaxHeap) Heapify(index int) {
	if heap.IsEmpty() {
		return
	}

	l := left(index)
	r := right(index)
	biggest := index

	if l < heap.Size() && (*heap)[l] > (*heap)[index] {
		biggest = l
	}

	if r < heap.Size() && (*heap)[r] > (*heap)[biggest] {
		biggest = r
	}

	if index != biggest {
		(*heap)[biggest], (*heap)[index] = (*heap)[index], (*heap)[biggest]
		heap.Heapify(biggest)
	}
}

// Size of max heap
func (heap *MaxHeap) Size() int {
	return len(*heap)
}

// IsEmpty max heap
func (heap *MaxHeap) IsEmpty() bool {
	return len(*heap) == 0
}

// ExtractMax remove max value from the max heap
func (heap *MaxHeap) ExtractMax() int {
	if len(*heap) == 0 {
		return -1
	}

	res := (*heap)[0]
	(*heap)[0] = (*heap)[heap.Size()-1]
	*heap = (*heap)[:heap.Size()-1]
	heap.Heapify(0)

	return res
}

// GetMax get max value in max heap
func (heap *MaxHeap) GetMax() int {
	if len(*heap) == 0 {
		return -1
	}

	return (*heap)[0]
}
