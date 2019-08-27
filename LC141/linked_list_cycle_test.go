/**
Given a linked list, determine if it has a cycle in it.

To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.



Example 1:

Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the second node.

Follow up:

Can you solve it using O(1) (i.e. constant) memory?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-cycle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package LC141

import (
	"testing"

	"github.com/klkyy2018/leetcode-go/utils"
)

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *utils.ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{
				utils.MkCycleList([]int{3, 2, 0, -4}, 1),
			},
			true,
		},
		{
			"test2",
			args{
				utils.MkCycleList([]int{1, 2}, 0),
			},
			true,
		},
		{
			"test3",
			args{
				utils.MkCycleList([]int{1, 2}, -1),
			},
			false,
		},
		{
			"test4",
			args{
				utils.MkCycleList([]int{1}, -1),
			},
			false,
		},
		{
			"empty list",
			args{
				utils.MkCycleList([]int{}, -1),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
