package main

import "fmt"

type LFUCache struct {
	minFre int
	kv map[int]*Node
	fv map[int]*DoubleList
}

type Node struct {
	key int
	val int
	fre int
	prev *Node
	next *Node
}

type DoubleList struct {
	head *Node
	tail *Node
	size int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		minFre: capacity,
		kv: map[int]*Node{},
		fv: map[int]*DoubleList{},
	}
}



func main()  {
	ppp := &Node{}
	fmt.Println(ppp)
}