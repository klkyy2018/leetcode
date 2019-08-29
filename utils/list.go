package utils

import (
	"fmt"
	"strconv"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func MkSingleList(input []int) *ListNode {
	if input == nil || len(input) == 0 {
		return nil
	}
	var iter, res *ListNode
	for _, item := range input {
		node := NewListNode(item)
		if iter != nil {
			iter.Next = node
		} else {
			res = node
		}
		iter = node
	}
	return res
}

func MkCycleList(input []int, pos int) *ListNode {
	if input == nil || len(input) == 0 {
		return nil
	}
	var iter, lintPos, res *ListNode
	for index, item := range input {
		node := NewListNode(item)
		if index == pos {
			lintPos = node
		} else if index == len(input)-1 {
			node.Next = lintPos
		}
		if iter != nil {
			iter.Next = node
		} else {
			res = node
		}
		iter = node
	}

	return res
}

func HasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func DetectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}

func SprintList(input *ListNode) string {
	res := ""
	postfix := "->"
	for {
		if input == nil {
			return res
		}
		if input.Next == nil {
			postfix = "."
		}
		res = res + strconv.Itoa(input.Val) + postfix
		input = input.Next
	}
}

func PrintList(input *ListNode) {
	postfix := "->"
	for {
		if input == nil {
			return
		}
		if input.Next == nil {
			postfix = ".\n"
		}
		fmt.Printf("%d%s", input.Val, postfix)
		input = input.Next
	}
}

func CompTwoList(l1 *ListNode, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	} else if l1 == nil && l2 != nil {
		return false
	} else if l1 != nil && l2 == nil {
		return false
	}
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}
