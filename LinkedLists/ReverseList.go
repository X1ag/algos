// https://leetcode.com/problems/palindrome-linked-list/

package main

type ListNode struct {
	Val int
	Next *ListNode 
} 

func middleNode(head *ListNode) *ListNode {
	s, f := head, head
	for f != nil && f.Next != nil {
			f = f.Next.Next
			s = s.Next
	}

	return s   
}

func reversedList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil {
			nextNode := head.Next
			head.Next = prev
			prev = head
			head = nextNode
	}
	return prev
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
			return true
	}
	headCopy := head

	rightSide := middleNode(headCopy)

	reversed := reversedList(rightSide)
	
	
	for reversed != nil {
			if head.Val != reversed.Val {
					return false
			}
			head = head.Next
			reversed = reversed.Next
	}
	return true
}
