package leetcode 

import (
	"fmt"
	"testing"
)

func L2S(list *ListNode) []int {
	return []int {
		list.Val,
		list.Next.Val,
		list.Next.Next.Val,
	}
}

func S2L(nums []int) *ListNode {
	return &ListNode{
		Val: nums[0],
		Next: &ListNode {
			Val: nums[1],
			Next: &ListNode {
				Val: nums[2],
				Next: nil,
			},
		},
	}
}

func compare(n []int, n2 []int) bool {
	return n[0] == n2[0] && n[1] == n2[1] && n[2] == n2[2]
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := [][]int{
		{2, 4, 3},
	}
	l2 := [][]int{
		{5, 6, 4},
	}
	target := [][]int{
		{7, 0, 8},
	}
	
	fmt.Printf("------------------------Leetcode Problem 2------------------------\n")
	for i:=0; i < len(l1); i++ {
		if compare(target[0], L2S(addTwoNumbers(S2L(l1[i]), S2L(l2[i])))) {
			t.Fatalf("case %d fails\n", i)
		}
	}
}