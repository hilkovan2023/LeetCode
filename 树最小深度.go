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
	if len(nums) == 0 {
		return nil
	}
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

func minnum(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func mindeep(root *node) int {
	min := -1
	if root == nil {
		min = 0
	} else if root.left == nil || root.right == nil {
		min = 1
	} else {
		min = 1 + minnum(mindeep(root.left), mindeep(root.right))
	}
	return min
}

func main() {
	root := newtree([]int{})
	printTree(root, "")
	res := mindeep(root)
	fmt.Println("最小高度：", res)
}
