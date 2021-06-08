package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

func preorder(root *node, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.val)
	preorder(root.left, res)
	preorder(root.right, res)
}

func main() {
	root := &node{10, nil, nil}
	node1 := &node{1, nil, nil}
	node2 := &node{2, nil, nil}
	node3 := &node{3, nil, nil}
	node4 := &node{4, nil, nil}
	node5 := &node{5, nil, nil}
	node6 := &node{6, nil, nil}
	root.left = node1
	root.right = node2
	root.left.left = node3
	root.left.right = node4
	root.right.right = node5
	root.right.right.right = node6
	res := []int{}
	preorder(root, &res)
	fmt.Println(res)
}
