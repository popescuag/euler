package list

import (
	"fmt"
	"math"
	"testing"
)

func AddTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := 0
	n2 := 0
	p := 0
	for l1.Next != nil {
		n1 += l1.Val * int(math.Pow10(p))
		p++
		l1 = l1.Next
	}
	n1 += l1.Val * int(math.Pow10(p))
	p = 0
	for l2.Next != nil {
		n2 += l2.Val * int(math.Pow10(p))
		p++
		l2 = l2.Next
	}
	n2 += l2.Val * int(math.Pow10(p))
	n := n1 + n2

	fmt.Println(n)

	currNode := ListNode{}
	head := &currNode
	tmp := &currNode
	for n >= 10 {
		tmp.Val = n % 10
		tmp.Next = &ListNode{}
		tmp = tmp.Next
		n /= 10
	}

	if n > 0 {
		tmp.Val = n
	}

	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n := 0
	currNode := ListNode{}
	head := &currNode
	tmp := &currNode
	l1Only, l2Only := false, false

	for {
		if l1Only {
			n = l1.Val + tmp.Val
		} else if l2Only {
			n = l2.Val + tmp.Val
		} else {
			n = l1.Val + l2.Val + tmp.Val
		}
		if n >= 10 {
			tmp.Val = n % 10
			tmp.Next = &ListNode{Val: 1}
		} else {
			tmp.Val = n
		}

		if l1.Next == nil && l2.Next == nil {
			break
		} else if n < 10 {
			tmp.Next = &ListNode{}
		}
		tmp = tmp.Next

		if l1.Next == nil {
			l2Only = true
			l2 = l2.Next
			//fmt.Println("l2only")
		} else if l2.Next == nil {
			l1Only = true
			l1 = l1.Next
			//fmt.Println("l1only")
		} else {
			l1 = l1.Next
			l2 = l2.Next
			//fmt.Println("l12only")
		}

	}

	return head
}

func TestSum(t *testing.T) {
	l1 := ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l1.printList()
	fmt.Println()
	l2 := ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	l2.printList()
	fmt.Println()
	sum := addTwoNumbers(&l1, &l2)
	sum.printList()
	t.Fail()
}
