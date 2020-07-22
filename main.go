package main

import (
	"fmt"

	"github.com/jadilet/generics/bst"
	"github.com/jadilet/generics/graph"
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

	// BST
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

	// MinHeap
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

	// MaxHeap
	m := heap.MaxHeap{}

	m.Insert(10)
	m.Insert(32)
	m.Insert(50)
	m.Insert(100)

	fmt.Println(m.ExtractMax())
	fmt.Println(m.ExtractMax())
	fmt.Println(m.ExtractMax())
	fmt.Println(m.ExtractMax())
	fmt.Println()

	// Graph implementation adjacency list
	gr := graph.Graph{
		"S": []graph.Vertex{{Name: "A", Weight: 7}, {Name: "B", Weight: 2}, {Name: "C", Weight: 3}},
		"A": []graph.Vertex{{Name: "D", Weight: 4}, {Name: "B", Weight: 3}, {Name: "S", Weight: 7}},
		"C": []graph.Vertex{{Name: "S", Weight: 3}, {Name: "L", Weight: 2}},
		"D": []graph.Vertex{{Name: "A", Weight: 4}, {Name: "B", Weight: 4}, {Name: "F", Weight: 5}},
		"F": []graph.Vertex{{Name: "H", Weight: 3}, {Name: "D", Weight: 5}},
		"H": []graph.Vertex{{Name: "F", Weight: 3}, {Name: "B", Weight: 1}, {Name: "G", Weight: 2}},
		"I": []graph.Vertex{{Name: "L", Weight: 4}, {Name: "J", Weight: 6}, {Name: "K", Weight: 4}},
		"J": []graph.Vertex{{Name: "K", Weight: 4}, {Name: "I", Weight: 6}, {Name: "L", Weight: 4}},
		"L": []graph.Vertex{{Name: "C", Weight: 2}, {Name: "I", Weight: 4}, {Name: "J", Weight: 4}},
		"K": []graph.Vertex{{Name: "I", Weight: 4}, {Name: "J", Weight: 4}, {Name: "E", Weight: 5}},
		"E": []graph.Vertex{{Name: "G", Weight: 2}, {Name: "K", Weight: 5}},
	}

	fmt.Println(gr.Dfs("S"))
	fmt.Println(gr.Bfs("S"))

}
