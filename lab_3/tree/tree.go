package tree

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

type Btree struct {
	Root *Node
}

func NewNode(val int) *Node {
	return &Node{val, nil, nil}
}

func Inorder(n *Node) {
	if n != nil {
		Inorder(n.right)
		fmt.Print(n.val, " ")
		Inorder(n.left)
	}
}

func Insert(root *Node, val int) *Node {
	if root == nil {
		return NewNode(val)
	}
	if val < root.val {
		root.left = Insert(root.left, val)
	} else {
		root.right = Insert(root.right, val)
	}
	return root
}
