/**
Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.

Note: Do not modify the linked list.



Example 1:

Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-cycle-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package LC142

import (
	"testing"

	"github.com/klkyy2018/leetcode-go/utils"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *utils.ListNode
	}
	tests := []struct {
		name string
		args args
		want *utils.ListNode
	}{
		{
			"test1",
			args{
				utils.MkCycleList([]int{3, 2, 0, -4}, 1),
			},
			utils.NewListNode(2),
		},
		{
			"test2",
			args{
				utils.MkCycleList([]int{1, 2}, 0),
			},
			utils.NewListNode(1),
		},
		{
			"test3",
			args{
				utils.MkCycleList([]int{1, 2}, -1),
			},
			nil,
		},
		{
			"test4",
			args{
				utils.MkCycleList([]int{1}, -1),
			},
			nil,
		},
		{
			"empty list",
			args{
				utils.MkCycleList([]int{}, -1),
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !utils.CompTwoList(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
