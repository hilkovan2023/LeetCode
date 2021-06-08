package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

func isSymmetric(left *node, right *node) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	} else if left.val != right.val {
		return false
	} else {
		fmt.Println("TTTTT")
		return isSymmetric(left.left, right.right) && isSymmetric(left.right, right.left)
	}
}

func printTree(root *node, head string) {
	const tag = "===//"
	head += tag
	if root == nil {
		fmt.Println(head)
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
	node2 := &node{1, nil, nil}
	node3 := &node{2, nil, nil}
	node4 := &node{2, nil, nil}
	node5 := &node{3, nil, nil}
	node6 := &node{3, nil, nil}
	node7 := &node{4, nil, nil}
	node8 := &node{4, nil, nil}
	node9 := &node{4, nil, nil}
	root.left = node1
	root.right = node2
	root.left.left = node3
	root.left.right = node5
	root.right.left = node6
	root.right.right = node4
	root.right.right.right = node7
	root.left.left.left = node8
	root.left.left.right = node9
	head := ""
	printTree(root, head)
	res := isSymmetric(root.left, root.right)
	fmt.Println(res)
}
