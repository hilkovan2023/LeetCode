package main

import "fmt"

type LinkNode struct {
	key  int
	val  int
	prev *LinkNode
	next *LinkNode
}

type LRUCache struct {
	cap  int
	head *LinkNode
	tail *LinkNode
	m    map[int]*LinkNode
}

func LRUConstructor(cap int) *LRUCache {
	head := &LinkNode{}
	tail := &LinkNode{}
	head.next = tail
	tail.prev = head
	return &LRUCache{
		cap,
		head,
		tail,
		make(map[int]*LinkNode),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.MoveToHead(v)
		return v.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok {
		v.val = value
		this.MoveToHead(v)
	} else {
		v := &LinkNode{
			key:  key,
			val:  value,
			prev: nil,
			next: nil,
		}
		if len(this.m) == this.cap {
			delete(this.m, this.tail.prev.key)
			this.RemoveNode(this.tail.prev)
		}
		this.AddNode(v)
		this.m[key] = v
	}
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddNode(node)
}

func (this *LRUCache) AddNode(node *LinkNode) {
	node.next = this.head.next
	node.next.prev = node
	this.head.next = node
	node.prev = this.head
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func main() {
	lru := LRUConstructor(3)
	fmt.Println(lru.Get(10))
	lru.Put(1, 11)
	lru.Put(2, 22)
	lru.Put(3, 33)
	lru.Put(4, 44)
	fmt.Println("m:", lru.m)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(2))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))
}
