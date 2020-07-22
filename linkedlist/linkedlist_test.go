package linkedlist

import "testing"

func TestInsert(t *testing.T) {
	list := LinkedList{}

	list.Insert(5)
	list.Insert(10)
	list.Insert(20)
	list.Insert(13)

	if list.head.data != 5 {
		t.Errorf("Insert incorrect head value, got: %v, expected %d", list.head.data, 5)
	}

	if list.tail.data != 13 {
		t.Errorf("Insert incorrect tail value, got: %v, expected %d", list.tail.data, 13)

	}
}

func TestDelete(t *testing.T) {
	list := LinkedList{}

	if list.head != nil {
		t.Errorf("Insert incorrect head value, got: %v, expected nil", list.head)
	}

	if list.tail != nil {
		t.Errorf("Insert incorrect tail value, got: %v, expected nil", list.tail)
	}

	list.Insert(10)
	list.Delete(10)

	if list.head != nil {
		t.Errorf("Insert incorrect head value, got: %v, expected nil", list.head)
	}

	if list.tail != nil {
		t.Errorf("Insert incorrect tail value, got: %v, expected nil", list.tail)
	}

	list.Insert(12)
	list.Insert(13)
	list.Insert(15)
	list.Insert(16)

	list.Delete(12)
	list.Delete(13)

	if list.head.data != 15 {
		t.Errorf("Delete incorrect head value, got: %v, expected %d", list.head.data, 15)
	}

	list.Delete(15)
	list.Delete(16)

	if list.head != nil {
		t.Errorf("Insert incorrect head value, got: %v, expected nil", list.head)
	}

	if list.tail != nil {
		t.Errorf("Insert incorrect tail value, got: %v, expected nil", list.tail)
	}

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)

	list.Delete(3)
	if list.tail.data != 2 {
		t.Errorf("Insert incorrect tail value, got: %v, expected %d", list.tail.data, 2)
	}

	list.Delete(2)
	
	if list.tail.data != 1 {
		t.Errorf("Insert incorrect tail value, got: %v, expected %d", list.tail.data, 1)
	}


	if list.head.data != 1 {
		t.Errorf("Insert incorrect head value, got: %v, expected %d", list.head.data, 1)
	}

	list.Delete(1)


	if list.head != nil {
		t.Errorf("Insert incorrect head value, got: %v, expected nil", list.head)
	}

	if list.tail != nil {
		t.Errorf("Insert incorrect tail value, got: %v, expected nil", list.tail)
	}

	list.Delete(-1)
	list.Delete(2)
	
	if list.head != nil {
		t.Errorf("Insert incorrect head value, got: %v, expected nil", list.head)
	}

	if list.tail != nil {
		t.Errorf("Insert incorrect tail value, got: %v, expected nil", list.tail)
	}
}
