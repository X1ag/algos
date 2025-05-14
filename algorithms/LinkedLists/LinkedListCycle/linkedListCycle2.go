// https://leetcode.com/problems/linked-list-cycle-ii/

package main

func detectCycle(head *ListNode) *ListNode {
    fast, slow := head, head
    for {
        if fast == nil || slow == nil || fast.Next == nil{
            return nil
        }

        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            break
        }
    }
    pos := head
    pos2 := slow
    
    for pos != pos2 {
        pos = pos.Next
        pos2 = pos2.Next
    }

    return pos
}