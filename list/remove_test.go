package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return head.Next
	}
	p := head
	p1 := head
	len := 1
	if n == 1 {
		for p.Next != nil {
			p = p.Next
			if p.Next != nil {
				p1 = p
			}
		}
		p1.Next = nil
		return head
	}
	dif := 0
	for {
		c := 0
		for c < n {
			c++
			len++
			dif++
			p = p.Next
			if p.Next == nil {
				break
			}
		}
		if p.Next == nil && (c == n-1 || len == n) {
			if p1 == head {
				return head.Next
			} else {
				p1.Next = p1.Next.Next
				return head
			}
		}
		if p.Next == nil {
			if len == n+1 {
				p1.Next = p1.Next.Next
				return head
			}
			for range n {
				p1 = p1.Next
			}
			p1.Next = p1.Next.Next
			return head
		}
		if dif > n {
			p1 = p
			dif = 0
		}
	}
}

func TestRemoveNth(t *testing.T) {
	lists := []ListNode{{Val: 1, Next: &ListNode{Val: 2}},
		{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}
	ns := []int{2, 2}
	results := []ListNode{{Val: 2}, {Val: 1, Next: &ListNode{Val: 3}}}

	for i := range lists {
		lists[i].printList()
		res := removeNthFromEnd(&lists[i], ns[i])
		res.printList()
		assert.Equal(t, &results[i], res)
	}
}
