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

func ringentrance(head *node) *node {
	fast := head
	slow := head
	fast = fast.next.next
	slow = slow.next
	for fast != slow {
		if fast.next == nil {
			return nil
		}
		fast = fast.next.next
		slow = slow.next
	}
	fast = head
	for fast != slow {
		fast = fast.next
		slow = slow.next
	}
	return fast
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

	node1 := &node{1, nil}
	node2 := &node{2, nil}
	node3 := &node{3, nil}
	node4 := &node{4, nil}
	node5 := &node{5, nil}
	node6 := &node{6, nil}
	node7 := &node{7, nil}
	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6
	node6.next = node7
	node7.next = node1
	rnode := ringentrance(node1)
	fmt.Println(rnode)

}
