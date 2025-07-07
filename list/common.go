package list

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) printList() {
	fmt.Print("[")
	for n.Next != nil {
		fmt.Printf("%v ", n.Val)
		n = n.Next
	}
	fmt.Printf("%v]\n", n.Val)
}
