package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

func invertTree(root *node) {
	if root == nil {
		return
	}
	root.left, root.right = root.right, root.left
	invertTree(root.left)
	invertTree(root.right)
}

func printTree(root *node, head string) {
	const tag = "===//"
	head += tag
	if root == nil {
		//fmt.Println(head)
		return
	}
	fmt.Println(head, root.val)
	printTree(root.left, head)
	printTree(root.right, head)
	head = head[4:]
}

func main() {
	root := &node{10, nil, nil}
	node1 := &node{1, nil, nil}
	node2 := &node{2, nil, nil}
	node3 := &node{3, nil, nil}
	node4 := &node{4, nil, nil}
	root.left = node1
	root.right = node2
	root.left.left = node3
	root.left.right = node4
	head := ""
	printTree(root, head)
	invertTree(root)
	printTree(root, head)
}
