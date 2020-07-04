package queue

import "testing"

var data  = "Google"

func TestQueue(t *testing.T) {
	var q Queue

	if len(q) != 0 {
		t.Errorf("Size of queue was incorrect,got %d, expected 0", len(q))
	}
}

func TestPush(t *testing.T) {
	var q Queue

	q.Push(data)

	if len(q) != 1 {
		t.Errorf("Push was incorrect size, got %d, expected: 1", len(q))
	}
}

func TestFront(t *testing.T) {
	var q Queue

	res := q.Front()

	if res != nil {
		t.Errorf("Front returned incorrec, got %v, expected nil", res)
	}

	q.Push(data)
	q.Push(data + " 1")
	q.Push(data + " 2")

	if q.Front() != data {
		t.Errorf("Front returned incorrect, got %s, expected %s", q.Front(), data)
	}
}

func TestPop(t *testing.T) {
	var q Queue

	q.Push(data)

	res := q.Pop()

	if res != data {
		t.Errorf("Pop returned incorrect, got %s, expected %s", res, data)
	}

	if len(q) != 0 {
		t.Errorf("Pop size was incorrect, got %d, expected 0", len(q))
	}

	res = q.Pop()

	if res != nil {
		t.Errorf("Pop returned incorrect, got %v, expected nil", res)
	}
}