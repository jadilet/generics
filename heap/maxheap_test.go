package heap

import (
	"testing"
)

func TestInsertMaxHeap(t *testing.T) {
	heap := &MaxHeap{}

	heap.Insert(10)

	if (*heap)[0] != 10 {
		t.Errorf("Insert not added value, got: %d, expected: 10", (*heap)[0])
	}
}

func TestSizeMaxHeap(t *testing.T) {
	heap := &MaxHeap{}

	heap.Insert(1)
	heap.Insert(2)

	if heap.Size() != 2 {
		t.Errorf("Size incorrect value returned, got: %d, expected: 2", heap.Size())
	}
}

func TestHeapifyMaxHeap(t *testing.T) {
	heap := &MaxHeap{}
	*heap = append(*heap, 3)
	*heap = append(*heap, 1)
	*heap = append(*heap, 2)

	heap.Heapify(0)

	//  MinHeap structure after heapifying.
	//      3
	//     / \
	//    1   2

	if heap.GetMax() != 3 {
		t.Errorf("Heapify not satisfied min heap requirement, got: %d, expected: 1", heap.GetMax())
	}

	if (*heap)[1] != 1 {
		t.Errorf("Heapify not satisfied min heap requirement, got: %d, expected: 1", (*heap)[1])
	}

	if (*heap)[2] != 2 {
		t.Errorf("Heapify not satisfied min heap requirement, got: %d, expected: 2", (*heap)[2])
	}
}

func TestGetMaxMaxHeap(t *testing.T) {
	heap := &MaxHeap{}

	heap.Insert(5)
	heap.Insert(10)
	heap.Insert(1)
	heap.Insert(3)

	if heap.GetMax() != 10 {
		t.Errorf("GetMin returned incorrect value, got: %d, expected: 1", heap.GetMax())
	}

	empty := &MaxHeap{}

	if empty.GetMax() != -1 {
		t.Errorf("GetMin returned incorrect value, got: %d, expected: -1", empty.GetMax())
	}
}

func TestIsEmptyMaxHeap(t *testing.T) {
	empty := &MaxHeap{}

	if !empty.IsEmpty() {
		t.Errorf("IsEmpty return incorrect value, got: %v, expected: true", empty.IsEmpty())
	}

	empty.Insert(1)

	if empty.IsEmpty() {
		t.Errorf("IsEmpty return incorrect value, got: %v, expected: false", empty.IsEmpty())
	}
}


func TestExtractMaxMaxHeap(t *testing.T) {
	heap := &MaxHeap{}

	heap.Insert(10)
	heap.Insert(5)
	heap.Insert(7)
	heap.Insert(16)
	heap.Insert(3)

	res := heap.ExtractMax() // 16

	if res != 16 {
		t.Errorf("ExtractMin returned incorrect value, got: %d, expected: 16", res)
	}

	res = heap.ExtractMax() // 10
	res = heap.ExtractMax() // 7
	res = heap.ExtractMax() // 5

	if res != 5 {
		t.Errorf("ExtractMin returned incorrect value, got: %d, expected: 5", res)
	}

	res = heap.ExtractMax()
	if res != 3 {
		t.Errorf("ExtractMin returned incorrect value, got: %d, expected: 3", res)
	}

	res = heap.ExtractMax()
	if res != -1 {
		t.Errorf("ExtractMin returned incorrect value, got: %d, expected: -1", res)
	}
}