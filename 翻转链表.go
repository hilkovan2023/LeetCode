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

//func reverselink(head *node) *node {
//	if head == nil || head.next == nil {
//		return head
//	}
//	fast := head.next
//	slow := head
//	slow.next = nil
//	for fast != nil {
//		pre := fast.next
//		fast.next = slow
//		slow = fast
//		fast = pre
//	}
//	return slow
//}

func reverselink(head *node) *node {
	pre := head
	cur := pre.next
	first := cur
	for cur != nil {
		next := cur.next
		cur.next = pre
		pre = cur
		cur = next
	}
	first.next = nil
	return pre
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
	head := buildlink([]int{12, 13, 1, 3, 2, 6, 8})
	printlink(head)
	rhead := reverselink(head)
	printlink(rhead)
}
