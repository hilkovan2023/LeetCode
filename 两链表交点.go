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

func intersection_of_links(head1 *node, head2 *node) *node {
	index1 := head1
	index2 := head2
	flag := 0 //防止无限循环
	for flag < 4 {
		if index1 == index2 {
			return index1
		}
		if index1.next == nil {
			index1 = head2
			flag++
			fmt.Println(flag)
		} else {
			index1 = index1.next
		}
		if index2.next == nil {
			index2 = head1
			flag++
			fmt.Println(flag)
		} else {
			index2 = index2.next
		}
	}
	return nil
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
	node3.next = node5
	node5.next = node6
	node4.next = node7
	snode := intersection_of_links(node1, node4)
	fmt.Println(snode)
	printlink(node1)
	printlink(node4)

	//fmt.Println("node4", node4)
	//fmt.Println("node7", node7)
	//*(node4.next) = node{1,nil}
	//fmt.Println("node4", node4)
	//fmt.Println("node7", node7)
}
