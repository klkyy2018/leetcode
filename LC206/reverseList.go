/**
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package LC206

import "github.com/klkyy2018/leetcode-go/utils"

func reverseList(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	foo := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return foo
}

func reverseListIter(head *utils.ListNode) *utils.ListNode {
	var prev *utils.ListNode
	curr := head
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}
