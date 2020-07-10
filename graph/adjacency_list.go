package graph

import (
	"github.com/jadilet/data-structure/queue"
	"github.com/jadilet/generics/stack"
)

// Graph data structure
type Graph map[interface{}][]interface{}

// AddEdge add edge
func (graph *Graph) AddEdge(src interface{}, dest interface{}) {
	(*graph)[src] = append((*graph)[src], dest)
}

// Dfs depth first search
func (graph *Graph) Dfs(vertex interface{}) []interface{} {
	var res []interface{}

	stack := stack.Stack{}
	visited := make(map[interface{}]bool)

	stack.Push(vertex)

	for !stack.IsEmpty() {
		ver := stack.Pop()

		if _, ok := visited[ver]; !ok {
			res = append(res, ver)
			visited[ver] = true
		}

		for _, v := range (*graph)[ver] {
			if _, ok := visited[v]; !ok {
				stack.Push(v)
			}
		}
	}

	return res
}

// Bfs breadth first search
func (graph *Graph) Bfs(vertex interface{}) []interface{} {
	var res []interface{}

	queue := queue.Queue{}
	visited := make(map[interface{}]bool)

	queue.Enqueue(vertex)

	for !queue.IsEmpty() {
		ver := queue.Dequeue()

		if _, ok := visited[ver]; !ok {
			res = append(res, ver)
			visited[ver] = true
		}

		for _, v := range (*graph)[ver] {
			if _, ok := visited[v]; !ok {
				queue.Enqueue(v)
			}
		}
	}

	return res
}
