package heap

// MinHeap data-structure
type MinHeap []int

// Insert to heap
func (heap *MinHeap) Insert(val int) {
	if len(*heap) == 0 {
		*heap = append(*heap, val)
		return
	}

	index := len(*heap)
	*heap = append(*heap, val)

	for index != 0 && (*heap)[parent(index)] > (*heap)[index] {
		(*heap)[parent(index)], (*heap)[index] = (*heap)[index], (*heap)[parent(index)]
		index = parent(index)
	}
}

// Size of heap
func (heap *MinHeap) Size() int {
	return len(*heap)
}

// ExtractMin delete min value
func (heap *MinHeap) ExtractMin() int {
	if len(*heap) == 0 {
		return -1
	}

	if len(*heap) == 1 {
		res := (*heap)[0]
		*heap = (*heap)[:len(*heap)-1]
		return res
	}

	res := (*heap)[0]
	(*heap)[0] = (*heap)[len(*heap)-1]
	*heap = (*heap)[:len(*heap)-1]
	heap.Heapify(0)

	return res
}

// Heapify updates the min heap from the index
func (heap *MinHeap) Heapify(index int) {
	l := left(index)
	r := right(index)
	smallest := index

	if l < len(*heap) && (*heap)[l] < (*heap)[index] {
		smallest = l
	}
	if r < len(*heap) && (*heap)[r] < (*heap)[smallest] {
		smallest = r
	}

	if smallest != index {
		(*heap)[index], (*heap)[smallest] = (*heap)[smallest], (*heap)[index]
		heap.Heapify(smallest)
	}
}

// GetMin return the smallest value in heap
// Time complexity O(1)
func (heap *MinHeap) GetMin() int {
	if heap.IsEmpty() {
		return -1
	}

	return (*heap)[0]
}

// IsEmpty heap
func (heap *MinHeap) IsEmpty() bool {
	return len(*heap) == 0
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}
