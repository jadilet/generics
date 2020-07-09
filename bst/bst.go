package bst

import (
	"fmt"

	"github.com/jadilet/generics/queue"
	"github.com/jadilet/generics/stack"
)

// Node of tree
type Node struct {
	val   int
	left  *Node
	right *Node
}

// BST binary search tree
type BST struct {
	root *Node
}

// Insert data to bst
func (bst *BST) Insert(val int) {
	if bst.root == nil {
		bst.root = &Node{val: val}
		return
	}

	insertHelper(bst.root, val)
}

func insertHelper(node *Node, val int) {
	if node == nil {
		return
	}

	if node.val >= val {
		if node.left == nil {
			node.left = &Node{val: val}
		} else {
			insertHelper(node.left, val)
		}
	} else {
		if node.right == nil {
			node.right = &Node{val: val}
		} else {
			insertHelper(node.right, val)
		}
	}
}

// Inorder traverse the BST
func (bst *BST) Inorder() {
	inorder(bst.root)
}

func inorder(root *Node) {
	if root == nil {
		return
	}

	if root.left != nil {
		inorder(root.left)
	}

	fmt.Println(root.val)

	if root.right != nil {
		inorder(root.right)
	}
}

// Dfs depth first search
// in order traverse
func (bst *BST) Dfs() {
	if bst.root == nil {
		return
	}

	var st stack.Stack	
	curr := bst.root

	for curr != nil || !st.IsEmpty() {
		if curr != nil {
			st.Push(curr)
			curr = curr.left
		} else {
			curr = st.Pop().(*Node)
			fmt.Println(curr.val)
			curr = curr.right
		}
	}
}

// Bfs breadth first search
func (bst *BST) Bfs() {

	if bst.root == nil {
		return
	}

	var queue queue.Queue
	queue.Push(bst.root)

	for !queue.IsEmpty() {
		node := queue.Front().(*Node)

		if node.left != nil {
			queue.Push(node.left)
		}

		if node.right != nil {
			queue.Push(node.right)
		}

		fmt.Println(node.val)
		queue.Pop()
	}

}
