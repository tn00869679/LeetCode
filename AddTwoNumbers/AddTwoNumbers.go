package AddTwoNumbers

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetAddTwoNumbersParams() (l1 *ListNode, l2 *ListNode) {
	l1 = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	return l1, l2
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0}
	current, carry := dummyHead, 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummyHead.Next
}

func AddTwoNumbersPrint(node *ListNode) {
	if node == nil {
		return
	}

	fmt.Println(node.Val)
	AddTwoNumbersPrint(node.Next)
}

func GetString(node *ListNode) string {
	var values []string
	for node != nil {
		values = append(values, strconv.Itoa(node.Val))
		node = node.Next
	}

	return strings.Join(values, " -> ")
}
