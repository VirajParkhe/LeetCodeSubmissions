/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    var prev  *ListNode
    temp := head
    for ; temp != nil ; {
        if temp.Val == val  && temp == head {
            head = head.Next
            temp = head
        }else if temp.Val == val {
            if temp.Next != nil {
                prev.Next = temp.Next
                temp = temp.Next
            }else {
                temp = nil
                prev.Next = nil
            }
        }else{
            temp = temp.Next
            if prev == nil {
                prev = head
            }else{
                prev = prev.Next     
            }
        }
        
    }
    return head
}