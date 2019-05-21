/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil{
        return nil
    }
    length := 1
    h := head
    for h.Next != nil {
        h = h.Next
        length++
    }
    h.Next = head
    k %= length
    for i := 0; i < length - k; i++ {
        h = h.Next
    }
    res := h.Next
    h.Next = nil
    return res
}