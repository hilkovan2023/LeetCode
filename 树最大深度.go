package main

import (
	"fmt"
)

type node struct {
	val   int
	left  *node
	right *node
}

func newtree(nums []int) *node {
	nodes := []*node{}
	for _, v := range nums {
		nodes = append(nodes, &node{val: v})
	}
	for i := 0; i < len(nums); i++ {
		if 2*i+1 < len(nums) {
			nodes[i].left = nodes[2*i+1]
		}
		if 2*i+2 < len(nums) {
			nodes[i].right = nodes[2*i+2]
		}
	}
	return nodes[0]
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

func maxnum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxdeep(root *node) int {
	if root == nil {
		return 0
	}
	max := 1 + maxnum(maxdeep(root.left), maxdeep(root.right))
	return max
}

func main() {
	root := newtree([]int{1, 2, 3, 4, 5, 6, 6})
	printTree(root, "")
	res := maxdeep(root)
	fmt.Println("树深度：", res)
}
