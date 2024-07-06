package test

import (
	"LeetCode/AddTwoNumbers"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1, l2 := AddTwoNumbers.GetAddTwoNumbersParams()
	result := AddTwoNumbers.AddTwoNumbers(l1, l2)
	// AddTwoNumbers.AddTwoNumbersPrint(result)
	expected := &AddTwoNumbers.ListNode{
		Val: 7,
		Next: &AddTwoNumbers.ListNode{
			Val: 0,
			Next: &AddTwoNumbers.ListNode{
				Val:  8,
				Next: nil,
			},
		},
	}

	resultToStr := AddTwoNumbers.GetString(result)
	expectedToStr := AddTwoNumbers.GetString(expected)

	if resultToStr != expectedToStr {
		t.Errorf("AddTwoNumbers(l1, l2) = %s; want %s", resultToStr, expectedToStr)
	}
}
