/*
* @Author: Ins
* @Date:   2019-04-01 13:02:43
* @Last Modified by:   Ins
* @Last Modified time: 2019-04-01 16:57:33
*/
package main

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    res := new(ListNode)
    res.Next = head
    pre, cur := res, res.Next
    for cur != nil {
    	p := pre.Next
    	i := k
    	for p != nil && i > 0 {
    		p = p.Next
    		i--
    	}
    	if i > 0 {
    		break
    	}
    	for cur.Next != p {
    		cur.Next.Next, pre.Next, cur.Next = pre.Next, cur.Next, cur.Next.Next
    		// t := cur.Next.Next
    		// cur.Next.Next = pre.Next
    		// pre.Next = cur.Next
    		// cur.Next = t
    	}
    	pre = cur
    	cur = pre.Next
    }
    return res.Next
}

