package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result = new(ListNode)
	var current = result
	var preCurrent *ListNode

	var buffer int
	var sum int

	for true {
		if l1 != nil || l2 != nil {
			if l1 != nil {
				sum += l1.Val
				l1 = l1.Next
			}
			if l2 != nil {
				sum += l2.Val
				l2 = l2.Next
			}
			sum += buffer
			if sum > 10 || sum == 10 {
				buffer = 1
				current.Val = sum - 10
			} else {
				buffer = 0
				current.Val = sum
			}
			sum = 0
			current.Next = new(ListNode)
			preCurrent = current
			current = current.Next
		} else {
			if buffer > 0 {
				current.Val = buffer
			} else {
				preCurrent.Next = nil
			}
			break
		}
	}
	return result
}
