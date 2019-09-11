package leetcode

import (
	"fmt"
	"testing"
)

func L2S(list *ListNode) []int {
	r := []int{}
	for list != nil {
		r = append(r, list.Val)
		list = list.Next
	}
	return r
}

func S2L(nums []int) *ListNode {
	head := &ListNode{}
	l := head
	for i := 0; i < len(nums); i++ {
		l.Next = &ListNode{Val: nums[i], Next: nil}
		l = l.Next
	}
	return head.Next
}

func compare(n1 []int, n2 []int) bool {
	if len(n1) != len(n2) {
		return false
	}
	r := true
	for i := 0; i < len(n1); i++ {
		r = r && (n1[i] == n2[i])
	}
	return r
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := [][]int{
		{2, 4, 3},
		{},
		{1, 2, 3, 4, 5},
	}
	l2 := [][]int{
		{5, 6, 4},
		{},
		{2, 2, 3, 4, 5},
	}
	target := [][]int{
		{7, 0, 8},
		{},
		{3, 4, 6, 8, 0, 1},
	}

	fmt.Printf("------------------------Leetcode Problem 2------------------------\n")
	for i := 0; i < len(l1); i++ {
		res := L2S(addTwoNumbers(S2L(l1[i]), S2L(l2[i])))
		if !compare(target[i], res) {
			t.Fatalf("case %d fails, %v\n", i, res)
		}
	}
}
