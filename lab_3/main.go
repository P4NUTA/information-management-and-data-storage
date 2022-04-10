package main

import (
	"fmt"
	"lab_3/tree"
)

func main() {
	t := &tree.Btree{}
	t.Root = tree.Insert(t.Root, 7)
	t.Root = tree.Insert(t.Root, 5)
	t.Root = tree.Insert(t.Root, 12)
	t.Root = tree.Insert(t.Root, 3)
	t.Root = tree.Insert(t.Root, 6)
	t.Root = tree.Insert(t.Root, 9)
	t.Root = tree.Insert(t.Root, 15)
	t.Root = tree.Insert(t.Root, 1)

	tree.Inorder(t.Root)
	fmt.Print("\n")
}
