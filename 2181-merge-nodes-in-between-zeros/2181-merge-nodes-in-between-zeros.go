/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    var prev, p1, p2, newHead *ListNode
    p1, p2 = head.Next, head.Next
    sum := 0
    for;p2 != nil; p2 = p2.Next{
        if p2.Val == 0 {
            p1.Val = sum
            sum = 0
            p1.Next = nil
            if prev != nil {
                prev.Next = p1
            }else{
                newHead = p1
            }
            prev = p1
            p1 = p2.Next
        }else {
            sum += p2.Val
        }
    }
    return newHead
}