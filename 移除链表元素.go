package main

import "fmt"

type node struct {
	val  int
	next *node
}

func makenodelist(num []int) *node {
	head := &node{}
	index := head
	for i := 0; i < len(num); i++ {
		cur := &node{val: num[i]}
		index.next = cur
		index = index.next
	}
	return head.next
}

func rmlele(head *node, target int) *node {
	if head == nil {
		return nil
	}
	prehead := &node{}
	prehead.next = head
	pre := prehead
	cur := head
	for cur != nil {
		if cur.val == target {
			//fmt.Println("===============", cur.val)
			cur = cur.next
			pre.next = cur
			//printlist(prehead)
		}
		pre = cur
		cur = cur.next
	}
	return prehead.next
}

func printlist(head *node) {
	index := head
	for index != nil {
		fmt.Print(index.val, "->")
		index = index.next
	}
	fmt.Println()
}

func main() {
	head := makenodelist([]int{1, 2, 3, 4, 5, 6})
	printlist(head)
	res := rmlele(head, 7)
	printlist(res)
}
