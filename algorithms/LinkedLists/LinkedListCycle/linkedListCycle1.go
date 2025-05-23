// https://leetcode.com/problems/linked-list-cycle/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
    fast, slow := head, head

    for {
        if fast == nil || slow == nil || fast.Next == nil{
            return false
        }

        fast = fast.Next.Next
        slow = slow.Next

        if fast == slow {
            return true
        }
    }
}