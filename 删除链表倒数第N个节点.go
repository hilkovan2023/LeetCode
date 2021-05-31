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

func dellastN(head *node, N int) *node {
	fast := head
	slow := head
	for i := 0; i < N; i++ {
		fast = fast.next
		if fast == nil {
			return head
		}
	}
	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
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
	shead := dellastN(head, 0)
	printlink(shead.next)
}
