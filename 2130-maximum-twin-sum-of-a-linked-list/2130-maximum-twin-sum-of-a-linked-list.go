/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var max int
var h *ListNode
func sum(tail *ListNode) {
    if tail == nil || h == tail || h == nil{
        return
    }
    sum(tail.Next)
    if ( h.Val + tail.Val ) > max {
        max = h.Val + tail.Val
    }
    h = h.Next
    return
}

func pairSum(head *ListNode) int {
    max = 0
    h = head
    sum(head.Next)
    return max
}