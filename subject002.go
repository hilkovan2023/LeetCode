package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 ==nil {
		return l1
	}

	l3 := new(ListNode)
	sum := l1.Val+l2.Val
	if sum < 10 {
		l3.Val = sum
		l3.Next = addTwoNumbers(l1.Next, l2.Next)
		return l3
	} else {
		l3.Val = sum%10
		temp := &ListNode{
			Val: 1,
			Next: nil,
		}
		l3.Next = addTwoNumbers(addTwoNumbers(l1.Next, l2.Next), temp)
		return l3
	}
}

func main() {
	l1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 6,
			Next: nil,
		},
	}

	l3 := addTwoNumbers(l1, l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}