/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    children := []*TreeNode{root, nil}
    prev := -1
    for ;; {
        sum :=0
        l := len(children)
        for _, val := range children {
            if val == nil {
                children = append(children, nil)
                break
            }
            sum+=val.Val
            if val.Left != nil {
                children = append(children, val.Left)
            }
            if val.Right != nil {
                children = append(children,val.Right )
            }
        }
        if sum != 0 {
            prev = sum
        }
        if sum == 0 {
            return prev
        }
       
        children = children[l:]      
    } 
    return -1
}