package main

import (
	"fmt"
)

//ListNode is the struct of linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	/*
		Test merge two sorted ll
		l1 := toLinkedList([]int{1, 2, 3, 4})
		l2 := toLinkedList([]int{1, 2})

		l3 := mergeTwoLists(l1, l2)
		l3.printLinkedList()
	*/

	/*
		// Test loop linked list
		l1 := toLoopLinkedList([]int{1, 2, 3, 4})
		l2 := toLinkedList([]int{1, 2, 3, 4})
		fmt.Println(hasCycle(l1))
		fmt.Println(hasCycle(l2))
		fmt.Println(hasCycle(nil))
	*/

}

func toLinkedList(arr []int) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: nil,
	}

	tail := dummyHead

	for _, n := range arr {
		node := &ListNode{
			Val:  n,
			Next: nil,
		}
		tail.Next = node
		tail = node
	}

	return dummyHead.Next
}

func toLoopLinkedList(arr []int) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: nil,
	}

	tail := dummyHead

	for _, n := range arr {
		node := &ListNode{
			Val:  n,
			Next: nil,
		}
		tail.Next = node
		tail = node
	}

	tail.Next = dummyHead.Next

	return dummyHead.Next
}

func (c *ListNode) printLinkedList() {
	step := 0
	curr := c
	for curr != nil {
		fmt.Println("This is ", step, " node ", "with value", curr.Val)
		curr = curr.Next
		step++
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	dummyHead := &ListNode{
		Val:  0,
		Next: nil,
	}

	tail := dummyHead

	curr1 := l1
	curr2 := l2
	for curr1 != nil && curr2 != nil {
		if curr1.Val <= curr2.Val {
			tail.Next = curr1
			tail = curr1
			curr1 = curr1.Next
		} else {
			tail.Next = curr2
			tail = curr2
			curr2 = curr2.Next
		}
	}

	if curr1 == nil {
		tail.Next = curr2
	} else {
		tail.Next = curr1
	}

	return dummyHead.Next

}

func hasCycle(head *ListNode) bool {

	//base case: head is null; head size is 1;
	if head == nil || head.Next == nil {
		return false
	}

	sp := head
	doubleSpeed := head.Next

	for doubleSpeed != sp && doubleSpeed.Next != nil {
		sp = sp.Next
		if doubleSpeed.Next.Next == nil {
			return false
		}
		doubleSpeed = doubleSpeed.Next.Next
	}

	if doubleSpeed == sp {
		return true
	}

	return false
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev, curr, next *ListNode
	prev = nil
	curr = head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func removeDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var first, curr *ListNode
	first = head
	curr = head.Next

	for curr != nil {
		for curr != nil && curr.Val == first.Val {
			curr = curr.Next
			first.Next = curr
		}

		first = curr

		if curr == nil {
			curr = curr.Next
		}

	}

	return head

}

func removeElements(head *ListNode, val int) *ListNode {

	var prev, curr *ListNode
	temHead := &ListNode{Val: 0, Next: head}
	prev = temHead
	curr = head
	for curr != nil {
		if curr.Val == val {
			curr = curr.Next
			prev.Next = curr
		} else {
			prev = curr
			curr = curr.Next
		}

	}

	return temHead.Next
}
