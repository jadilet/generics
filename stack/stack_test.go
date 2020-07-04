package stack

import "testing"

var data = "Google"

func TestPush(t *testing.T) {
	var st Stack

	st.Push(data)

	if len(st) != 1 {
		t.Errorf("Push was incorrect size, got %d, want 1", len(st))
	}
}

func TestTop(t *testing.T) {
	var st Stack

	if st.Top() != nil {
		t.Errorf("Top returned incorrect data, got %v, expected nil", st.Top())
	}

	st.Push(data)

	if st.Top() != data {
		t.Errorf("Top was incorrect, got %s, want Google", st.Top())
	}
}

func TestPop(t *testing.T) {
	var s Stack

	if s.Pop() !=  nil {
		t.Errorf("Pop returned incorrect, got %v, expected nil", s.Pop())
	}

	s.Push(data)
	s.Pop()

	if len(s) != 0 {
		t.Errorf("Pop was incorrect, got %d, expected 0", len(s))
	}
}

func TestIsEmpty(t *testing.T) {
	var s Stack

	s.Push("test")
	if s.IsEmpty() != false {
		t.Errorf("IsEmpty was correct, expected true, got %v", s.IsEmpty())
	}
}

func TestStack(t *testing.T) {
	var s Stack

	if len(s) != 0 {
		t.Errorf("Size of stack was incorrect, expected 0, got %d", len(s))
	}
}

func BenchmarkPop(b *testing.B) {
	var s Stack

	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func BenchmarkPush(b *testing.B) {
	var s Stack

	for i := 0; i < b.N; i++ {
		s.Push(data)
	}
}

func BenchmarkTop(b *testing.B) {
	var s Stack
	s.Push(data)

	for i := 0; i < b.N; i++ {
		s.Top()
	}
}