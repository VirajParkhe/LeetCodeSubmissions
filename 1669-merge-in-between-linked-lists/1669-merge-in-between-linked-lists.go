/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    head :=  list1
    var temp1, temp2 *ListNode
    count := 0
    for ;head != nil;head = head.Next{
        if count +1 == a {
            temp1 = head
        }
        if count == b {
            temp2 = head.Next
        }
        count++
    }
    temp1.Next=list2
    head2 := list2
    for ;head2.Next !=nil;  {
        head2= head2.Next
    }
    head2.Next = temp2
    return list1
}