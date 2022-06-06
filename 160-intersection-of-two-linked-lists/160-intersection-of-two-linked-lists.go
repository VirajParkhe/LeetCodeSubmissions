/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    m := map[*ListNode]struct{}{}
    for ;headA!=nil; headA=headA.Next{
        m[headA] = struct{}{}
    }
    for ; headB != nil;headB=headB.Next {
        if _, ok := m[headB]; ok {
            return headB
        }
    }
    return nil
}