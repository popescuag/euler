package list

import (
	"fmt"
	"testing"
)

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmpNode := head
	arr := []*ListNode{head}
	for tmpNode.Next != nil {
		arr = append(arr, tmpNode)
		tmpNode = tmpNode.Next
	}
	arr = append(arr, tmpNode)
	reverseHead := tmpNode

	for i := len(arr) - 1; i > 0; i-- {
		arr[i].Next = arr[i-1]
	}
	arr[0].Next = nil

	return reverseHead
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	reverseHead := reverseList2(head.Next)

	head.Next.Next = head
	head.Next = nil

	return reverseHead
}

func TestReverse(t *testing.T) {
	list := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	list.printList()
	reversed := reverseList(&list)
	reversed.printList()
	t.Fail()
}

func TestReverse2(t *testing.T) {
	list := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	//list.printList()
	reversed := reverseList2(&list)
	reversed.printList()

	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr[:3])
	fmt.Println(arr[3:6])

	t.Fail()
}
