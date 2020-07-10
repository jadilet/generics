package graph

import (
	"reflect"
	"testing"
)

func TestAddEdge(t *testing.T) {
	graph := &Graph{}

	//        0 
   //       / \  \
   //      1---2  3
   //          |
   //          4
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(0, 3)
	graph.AddEdge(2, 4)

	// Edge 0
	var expected[]interface{} = []interface{}{1,2,3}

	if !reflect.DeepEqual(expected, (*graph)[0]) {
		t.Errorf("AddEdge incorrect graph vertices, got: %v, expected: %v", 
		(*graph)[0], expected)
	}
}

func TestDfs(t *testing.T) {
	graph := &Graph{}

	//        0 
   //       / \  \
   //      1---2  3
   //          |
   //          4
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(0, 3)
	graph.AddEdge(2, 4)
	
	got := graph.Dfs(0)
	var expected[]interface{} = []interface{}{0, 3, 2, 4, 1}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Dfs incorrect graph traverse, got: %v, expected: %v", 
		got, expected)
	}

}

func TestBfs(t *testing.T) {
	graph := &Graph{}

	//        0 
   //       / \  \
   //      1---2  3
   //          |
   //          4
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(0, 3)
	graph.AddEdge(2, 4)
	
	got := graph.Bfs(0)
	
	var expected[]interface{} = []interface{}{0, 1, 2, 3, 4}
	
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Bfs incorrect graph traverse, got: %v, expected: %v", 
		got, expected)
	}
}