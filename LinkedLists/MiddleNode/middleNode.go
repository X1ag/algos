// https://leetcode.com/problems/middle-of-the-linked-list/

package main

type ListNode struct {
	Val  int
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