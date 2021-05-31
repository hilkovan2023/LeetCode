package main

import "fmt"

type node struct {
	val  int
	next *node
}

func buildlink(num []int) *node {
	head := &node{}
	cur := head
	for i := 0; i < len(num); i++ {
		node := &node{val: num[i]}
		cur.next = node
		cur = cur.next
	}
	return head
}

func switch2node(head *node) *node {
	pre := head
	for pre.next != nil && pre.next.next != nil {
		index1 := pre.next
		index2 := pre.next.next
		//fmt.Println(index1.val, index2.val)

		pre.next = index2
		pre = index2.next
		index2.next = index1
		index1.next = pre
		pre = index1
	}
	return head
}

func printlink(head *node) {
	cur := head
	for cur != nil {
		fmt.Printf("%v -> ", cur.val)
		cur = cur.next
	}
	fmt.Println()
}

func main() {
	head := buildlink([]int{1, 2, 3, 4, 5, 6})
	printlink(head.next)
	shead := switch2node(head)
	printlink(shead.next)
}
