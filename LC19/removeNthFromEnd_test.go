package LC19

import (
	"fmt"
	"github.com/klkyy2018/leetcode-go/utils"
	"strconv"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T)  {
	inputs := []struct {
		input []int
		n int
		wanted *utils.ListNode
	}{
		{
			[]int{1},
			1,
			utils.MkSingleList([]int{}),
		},
		{
			[]int{1,2},
			1,
			utils.MkSingleList([]int{1}),
		},
		{
			[]int{1,2},
			2,
			utils.MkSingleList([]int{2}),
		},
		{
			[]int{1,2,3},
			1,
			utils.MkSingleList([]int{1,2}),
		},
		{
			[]int{1,2,3},
			2,
			utils.MkSingleList([]int{1,3}),
		},
		{
			[]int{1,2,3},
			3,
			utils.MkSingleList([]int{2,3}),
		},
		{
			[]int{},
			0,
			utils.MkSingleList([]int{}),
		},
	}
	for _, elem := range inputs {
		inputList := utils.MkSingleList(elem.input)
		hint := fmt.Sprintf("input is " + utils.SprintList(inputList) + " n is " + strconv.Itoa(elem.n))
		outputList := removeNthFromEnd(inputList, elem.n)
		if !utils.CompTwoList(elem.wanted, outputList) {
			t.Errorf("output is %s, wanted is %s, %s", utils.SprintList(outputList), utils.SprintList(elem.wanted), hint)
		}
	}
}
