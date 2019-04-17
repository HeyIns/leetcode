
//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}
 
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 || len(lists) == 1 && lists[0] == nil {
        return nil
    }
    //setup
    heap := []*ListNode{}
    for _, v := range(lists) {
        if v != nil {
            heap = append(heap, v)
            i := len(heap) - 1
            j := (i + 1) / 2 - 1
            for i > 0 && heap[i].Val < heap[j].Val {
                heap[i], heap[j] = heap[j], heap[i]
                i = j
                j = (i + 1) / 2 - 1
            }
        }
    }
    
    res := new(ListNode)
    result := res
    for len(heap) > 0 {
        res.Next = heap[0]
        res = res.Next
        if heap[0].Next == nil {
            heap[0] = heap[len(heap) - 1]
            heap = heap[:len(heap) - 1]
        } else {
            heap[0] = heap[0].Next
        }
        i := 0
        for i < len(heap) {
            child := i
            if 2 * i + 2 < len(heap) {
                if heap[2 * i + 1].Val <= heap[2 * i + 2].Val {
                    child = 2 * i + 1
                } else {
                    child = 2 * i + 2
                }
            } else if 2 * i + 1 < len(heap) {
                child = 2 * i + 1
            }
            if heap[i].Val > heap[child].Val {
                heap[i], heap[child] = heap[child], heap[i]
                i = child
            } else {
                break
            }
        }
    }
    return result.Next
}