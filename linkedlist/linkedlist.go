package linkedlist

// Node of linkedlist
type Node struct {
	data interface{}
	next *Node
}

// LinkedList data structure
type LinkedList struct {
	head *Node
	tail *Node
}

// Insert data to back of linkedlist
// O(1) time complexity
func (list *LinkedList) Insert(data interface{}) {
	if list.head == nil {
		list.head = &Node{data: data}
		list.tail = list.head
		return
	}

	list.tail.next = &Node{data: data}
	list.tail = list.tail.next
}

// Delete from linkedlist
// O(n) time complexity
func (list *LinkedList) Delete(data interface{}) {
	list.head = deleteHelper(list.head, data)
	list.updateTail()
}

func deleteHelper(node *Node, data interface{}) *Node {
	if node == nil {
		return nil
	}

	if node.data == data {
		return node.next
	}

	node.next = deleteHelper(node.next, data)

	return node
}

func (list *LinkedList) updateTail() {
	tmp := list.head

	for tmp != nil && tmp.next != nil {
		tmp = tmp.next
	}

	list.tail = tmp
}
