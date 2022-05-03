/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    var p1, head2 *ListNode
    p2 := head.Next
    sum := 0
    for ;p2 != nil; {
        if p2.Val == 0 {
            temp :=&ListNode{sum, nil}
            if p1 == nil {
                p1 = temp
                head2 = p1
            }else{
               p1.Next = temp
               p1= temp
            }
            p2 = p2.Next
            sum =0
        }else{
            sum+=p2.Val
            p2 = p2.Next
        }
    }
    return head2
}