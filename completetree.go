package main

import "fmt"

type node struct {
	val int
	left *node
	right *node
}

func buildtree(num []int) *node {
	if len(num)==0 {
		return nil
	}
	list := []*node{}
	index := 0
	root := &node{val: num[index]}
	list = append(list, root)
	index++

	for  {
		if index < len(num) {
			list[0].left = &node{val:num[index]}
			fmt.Println(list[0].left)
			list = append(list, list[0].left)
			index++
		}
		if index < len(num) {
			list[0].right = &node{val:num[index]}
			fmt.Println(list[0].right)
			list = append(list, list[0].right)
			index++
		}
		list = list[1:]
		if index == len(num) {
			break
		}
	}
	return root
}

var res []int

func treeprint(root *node) []int {
	if root == nil {
		return nil
	}
	list := []*node{}
	list = append(list, root)
	printoneline(list)
	return res
}

func printoneline(list []*node) {
	if len(list) == 0 {
		return
	}
	n := len(list)
	for i:=0; i<n; i++ {
		res = append(res, list[i].val)
		fmt.Println("====", res)
		if list[i].left != nil {
			list = append(list, list[i].left)
		}
		if list[i].right != nil {
			list = append(list, list[i].right)
		}
	}
	printoneline(list[n:])
}

func main() {
	root := buildtree([]int{1,2,3,4,5,6,7,8,9})
	num := treeprint(root)
	fmt.Println(num)
}
