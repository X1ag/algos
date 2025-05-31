// https://leetcode.com/problems/palindrome-linked-list/

package main

type ListNode struct {
	Val int
	Next *ListNode 
} 

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		 fast = fast.Next.Next
		 slow = slow.Next
	}

	return slow
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

 right := middleNode(head)
 reversed := reversedList(right)

 for reversed != nil {
		 if reversed.Val != head.Val {
				 return false
		 }
		 reversed = reversed.Next
		 head = head.Next
 }
 return true
}
