//get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
//addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
//addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
//addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
//deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

package main

import "fmt"

type node struct {
	val  int
	next *node
}

type link struct {
	size int
	head *node
}

func newlink() *link {
	head := &node{}
	mylink := &link{size: 0, head: head}
	return mylink
}

func (mylink *link) addAtHead(val int) {
	node := &node{val: val}
	node.next = mylink.head.next
	mylink.head.next = node
	mylink.size++
}

func (mylink *link) addAtTail(val int) {
	cur := mylink.head
	for i := 0; i < mylink.size; i++ {
		cur = cur.next
	}
	node := &node{val: val}
	cur.next = node
	mylink.size++
}

func (mylink *link) addAtIndex(index int, val int) {
	if index <= 0 {
		mylink.addAtHead(val)
	} else if index > mylink.size {
		return
	} else {
		cur := mylink.head
		for i := 0; i < index; i++ {
			cur = cur.next
		}
		node := &node{val: val}
		node.next = cur.next
		cur.next = node
		mylink.size++
	}
}

func (mylink *link) deleteAtIndex(index int) {
	if index < 0 || index >= mylink.size {
		return
	}
	cur := mylink.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.next = cur.next.next
	mylink.size--
}

func (mylink *link) get(index int) int {
	if index < 0 || index >= mylink.size {
		return -1
	}
	cur := mylink.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.next.val

}

func printlink(mylink *link) {
	if mylink.size == 0 {
		fmt.Println("empty link!")
		return
	}
	cur := mylink.head.next
	for i := 0; i < mylink.size; i++ {
		fmt.Print(cur.val, "->")
		cur = cur.next
	}
	fmt.Println()
}

func main() {
	mylink := newlink()
	mylink.addAtHead(1)
	mylink.addAtTail(3)
	mylink.addAtTail(4)
	mylink.addAtIndex(1, 2)
	mylink.addAtIndex(5, 12)
	fmt.Println("get", mylink.get(0))
	printlink(mylink)
	mylink.deleteAtIndex(1)
	printlink(mylink)
	fmt.Println("get", mylink.get(1))

}
