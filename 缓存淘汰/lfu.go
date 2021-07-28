package main

import "fmt"

type Node struct {
	key  int
	val  int
	fre  int
	prev *Node
	next *Node
}

type DoubleList struct {
	size int
	head *Node
	tail *Node
}

type LFUCache struct {
	cap  int
	minF int
	kv   map[int]*Node
	fv   map[int]*DoubleList
}

func LFUConstructor(cap int) *LFUCache {
	kvHash := make(map[int]*Node)
	fvHash := make(map[int]*DoubleList)
	return &LFUCache{
		cap:  cap,
		minF: 0,
		kv:   kvHash,
		fv:   fvHash,
	}
}

func LinkConstructor() *DoubleList {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &DoubleList{
		size: 0,
		head: head,
		tail: tail,
	}
}

func (this *LFUCache) Get(key int) int {
	if curNode, ok := this.kv[key]; ok {
		this.increaseKeyFre(key)
		return curNode.val
	}
	return -1
}

func (this *LFUCache) Put(key int, val int) {
	if this.cap <= 0 {
		return
	}

	if this.Get(key) == -1 {
		if len(this.kv) >= this.cap {
			this.removeMinFreKey()
		}

		newNode := &Node{key: key, val: val, fre: 1, prev: nil, next: nil}
		this.kv[key] = newNode
		if lastLink, ok := this.fv[1]; ok {
			lastLink.AddHead(newNode)
		} else {
			lastLink := LinkConstructor()
			lastLink.AddHead(newNode)
			this.fv[1] = lastLink
		}
		this.minF = 1

	} else {
		this.kv[key].val = val
		this.increaseKeyFre(key)
	}
}

func (dl *DoubleList) RemoveNode(node *Node) {
	if dl.size == 0 {
		return
	}
	dl.size--
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (dl *DoubleList) AddHead(node *Node) {
	dl.size++
	node.next = dl.head.next
	dl.head.next.prev = node
	dl.head.next = node
	node.prev = dl.head
}

func (lfu *LFUCache) increaseKeyFre(key int) {
	curNode := lfu.kv[key]
	curLink := lfu.fv[curNode.fre]
	if curLink.size == 1 {
		delete(lfu.fv, curNode.fre)
	} else {
		curLink.RemoveNode(curNode)
	}
	curNode.fre++
	if newLink, ok := lfu.fv[curNode.fre]; ok {
		newLink.AddHead(curNode)
	} else {
		newLink := LinkConstructor()
		newLink.AddHead(curNode)
		lfu.fv[curNode.fre] = newLink
	}
}

func (lfu *LFUCache) removeMinFreKey() {
	lastLink := lfu.fv[lfu.minF]
	lastNode := lastLink.tail.prev
	if lastLink.size == 1 {
		delete(lfu.fv, lfu.minF)
	} else {
		lastLink.RemoveNode(lastNode)
	}
	delete(lfu.kv, lastNode.key)
}

func main() {
	lfu := LFUConstructor(3)
	lfu.Put(1, 11)
	lfu.Put(2, 22)
	lfu.Put(3, 33)
	lfu.Put(4, 44)
	fmt.Println(lfu.kv)
	fmt.Println(lfu.fv)
	fmt.Println(lfu.Get(1))
	fmt.Println(lfu.kv)
	fmt.Println(lfu.fv)
	fmt.Println(lfu.Get(2))
	fmt.Println(lfu.fv)
	fmt.Println(lfu.Get(3))
	fmt.Println(lfu.fv)
	fmt.Println(lfu.Get(4))
	fmt.Println(lfu.fv)
}
