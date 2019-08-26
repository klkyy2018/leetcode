/**
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
package LC19

import "github.com/klkyy2018/leetcode-go/utils"

func removeNthFromEnd(head *utils.ListNode, n int) *utils.ListNode  {
	if head == nil {
		return nil
	}
	p1 := head
	p2 := head
	for i:=0; i<=n; i++ {
		p2 = p2.Next
		if p2 == nil && i != n {
			return head.Next
		}
	}
	for p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	if p1.Next != nil {
		p1.Next = p1.Next.Next
	} else {
		p1.Next = nil
	}
	return head
}