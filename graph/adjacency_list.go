package graph

import (
	"github.com/jadilet/generics/queue"
	"github.com/jadilet/generics/stack"
)

// Graph data structure
type Graph map[interface{}][]Vertex

// Vertex graph
type Vertex struct {
	Name   interface{}
	Weight interface{}
}

// AddEdge add edge
func (graph *Graph) AddEdge(src interface{}, dest Vertex) {
	(*graph)[src] = append((*graph)[src], dest)
}

// Dfs depth first search
func (graph *Graph) Dfs(edge interface{}) []interface{} {
	var res []interface{}

	stack := stack.Stack{}
	visited := make(map[interface{}]bool)

	stack.Push(edge)

	for !stack.IsEmpty() {
		ver := stack.Pop()

		if _, ok := visited[ver]; !ok {
			res = append(res, ver)
			visited[ver] = true
		}

		for _, v := range (*graph)[ver] {
			if _, ok := visited[v.Name]; !ok {
				stack.Push(v.Name)
			}
		}
	}

	return res
}

// Bfs breadth first search
func (graph *Graph) Bfs(edge interface{}) []interface{} {
	var res []interface{}

	queue := queue.Queue{}
	visited := make(map[interface{}]bool)

	queue.Push(edge)

	for !queue.IsEmpty() {
		ver := queue.Pop()

		if _, ok := visited[ver]; !ok {
			res = append(res, ver)
			visited[ver] = true
		}

		for _, v := range (*graph)[ver] {
			if _, ok := visited[v.Name]; !ok {
				queue.Push(v.Name)
			}
		}
	}

	return res
}
