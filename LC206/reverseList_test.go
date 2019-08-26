package LC206

import (
	"fmt"
	"github.com/klkyy2018/leetcode/utils"
	"testing"
)

func TestReverseList(t *testing.T)  {
	inputs := []struct {
		input []int
		wanted *utils.ListNode
	}{
		{
			[]int{1,2,3,4,5},
			utils.MkSingleList([]int{5,4,3,2,1}),
		},
		{
			[]int{},
			utils.MkSingleList([]int{}),
		},
		{
			[]int{1},
			utils.MkSingleList([]int{1}),
		},
	}
	for _, input := range inputs {
		inputList := utils.MkSingleList(input.input)
		hint := fmt.Sprintf("input is " + utils.SprintList(inputList))
		output := reverseList(inputList)
		if !utils.CompTwoList(output, input.wanted) {
			t.Errorf("output: %s, wanted %s. %s", utils.SprintList(output), utils.SprintList(input.wanted), hint)
		}
	}
}

func TestReverseListIter(t *testing.T)  {
	inputs := []struct {
		input []int
		wanted *utils.ListNode
	}{
		{
			[]int{1,2,3,4,5},
			utils.MkSingleList([]int{5,4,3,2,1}),
		},
		{
			[]int{},
			utils.MkSingleList([]int{}),
		},
		{
			[]int{1},
			utils.MkSingleList([]int{1}),
		},
	}
	for _, input := range inputs {
		inputList := utils.MkSingleList(input.input)
		hint := fmt.Sprintf("input is " + utils.SprintList(inputList))
		output := reverseListIter(inputList)
		if !utils.CompTwoList(output, input.wanted) {
			t.Errorf("output: %s, wanted %s. %s", utils.SprintList(output), utils.SprintList(input.wanted), hint)
		}
	}
}
