package main

import (
	"fmt"

	"github.com/jadilet/generics/stack"
	"github.com/jadilet/generics/queue"
	
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
}
