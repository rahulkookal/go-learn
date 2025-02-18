package bst

import "fmt"

type Node struct {
	Value int64
	left  *Node
	right *Node
}

func (node *Node) Insert(i int64) {
	if i < node.Value {
		if node.left == nil {
			node.left = &Node{
				Value: i,
			}
		} else {
			node.left.Insert(i)
		}

	} else {
		if node.right == nil {
			node.right = &Node{Value: i}
		} else {
			node.right.Insert(i)
		}
	}
}

func (n *Node) InOrderTraversal() {
	if n == nil {
		return
	}
	n.left.InOrderTraversal()
	fmt.Print(n.Value, "->")
	n.right.InOrderTraversal()
}

func (n *Node) PreOrder() {
	if n == nil {
		return
	}
	fmt.Print(n.Value, "->")
	n.left.PreOrder()
	n.right.PreOrder()
}
