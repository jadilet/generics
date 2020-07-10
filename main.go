package main

import (
	"fmt"

	"github.com/jadilet/generics/bst"
	"github.com/jadilet/generics/heap"
	"github.com/jadilet/generics/queue"
	"github.com/jadilet/generics/stack"
)

func main() {

	var s stack.Stack
	var q queue.Queue

	s.Push("Google")
	q.Push("Test 1")
	q.Push("Test 2")
	q.Push("Test 3")

	fmt.Println(s.Top())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	bst := bst.BST{}

	bst.Insert(10)
	bst.Insert(9)
	bst.Insert(12)
	bst.Insert(11)
	bst.Insert(13)
	bst.Insert(8)

	fmt.Print(" DFS ")
	bst.Dfs()
	fmt.Println("---")

	fmt.Println("BFS")
	bst.Bfs()
	fmt.Println("---")

	h := heap.MinHeap{}

	h.Insert(10)
	h.Insert(5)
	h.Insert(9)
	h.Insert(2)

	fmt.Println(h.ExtractMin()) // 2
	fmt.Println(h.ExtractMin()) // 5
	fmt.Println(h.ExtractMin()) // 9
	fmt.Println(h.ExtractMin()) // 10
	fmt.Println(h.ExtractMin())

	m := heap.MaxHeap{}

	m.Insert(10)
	m.Insert(32)
	m.Insert(50)
	m.Insert(100)

	fmt.Println(m.ExtractMax())
	fmt.Println(m.ExtractMax())
	fmt.Println(m.ExtractMax())
	fmt.Println(m.ExtractMax())

}
