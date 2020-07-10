package heap

import (
	"testing"
)

func TestInsert(t *testing.T) {
	heap := &MinHeap{}

	heap.Insert(10)

	if (*heap)[0] != 10 {
		t.Errorf("Insert not added value, got: %d, expected: 10", (*heap)[0])
	}
}

func TestSize(t *testing.T) {
	heap := &MinHeap{}

	heap.Insert(1)
	heap.Insert(2)

	if heap.Size() != 2 {
		t.Errorf("Size incorrect value returned, got: %d, expected: 2", heap.Size())
	}
}

func TestHeapify(t *testing.T) {
	heap := &MinHeap{}
	*heap = append(*heap, 3)
	*heap = append(*heap, 1)
	*heap = append(*heap, 2)

	heap.Heapify(0)

	//  MinHeap structure after heapifying.
	//      1
	//     / \
	//    3   2

	if heap.GetMin() != 1 {
		t.Errorf("Heapify not satisfied min heap requirement, got: %d, expected: 1", heap.GetMin())
	}

	if (*heap)[1] != 3 {
		t.Errorf("Heapify not satisfied min heap requirement, got: %d, expected: 3", (*heap)[1])
	}

	if (*heap)[2] != 2 {
		t.Errorf("Heapify not satisfied min heap requirement, got: %d, expected: 2", (*heap)[2])
	}
}

func TestGetMin(t *testing.T) {
	heap := &MinHeap{}

	heap.Insert(5)
	heap.Insert(10)
	heap.Insert(1)
	heap.Insert(3)

	if heap.GetMin() != 1 {
		t.Errorf("GetMin returned incorrect value, got: %d, expected: 1", heap.GetMin())
	}

	empty := &MinHeap{}

	if empty.GetMin() != -1 {
		t.Errorf("GetMin returned incorrect value, got: %d, expected: -1", empty.GetMin())
	}
}

func TestIsEmpty(t *testing.T) {
	empty := &MinHeap{}

	if !empty.IsEmpty() {
		t.Errorf("IsEmpty return incorrect value, got: %v, expected: true", empty.IsEmpty())
	}

	empty.Insert(1)

	if empty.IsEmpty() {
		t.Errorf("IsEmpty return incorrect value, got: %v, expected: false", empty.IsEmpty())
	}
}

func TestExtractMin(t *testing.T) {
	heap := &MinHeap{}

	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(7)
	heap.Insert(16)
	heap.Insert(3)

	res := heap.ExtractMin() // 3

	if res != 3 {
		t.Errorf("ExtractMin returned incorrect value, got: %d, expected: 3", res)
	}

	res = heap.ExtractMin() // 5
	res = heap.ExtractMin() // 7
	res = heap.ExtractMin() // 10

	if res != 10 {
		t.Errorf("ExtractMin returned incorrect value, got: %d, expected: 10", res)
	}

	res = heap.ExtractMin()
	if res != 16 {
		t.Errorf("ExtractMin returned incorrect value, got: %d, expected: 16", res)
	}

	res = heap.ExtractMin()
	if res != -1 {
		t.Errorf("ExtractMin returned incorrect value, got: %d, expected: -1", res)
	}
}
