/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root ==  nil {
        return &TreeNode{val, nil,nil}
    }
    if root.Val > val {
        if root.Left == nil {
            root.Left = &TreeNode{val, nil,nil}
            return root
        }else{
            root.Left = insertIntoBST(root.Left, val)   
        }
    }
    if root.Val < val {
        if root.Right == nil {
            root.Right = &TreeNode{val,nil,nil}
            return root
        }else{
            root.Right =insertIntoBST(root.Right, val)
        }
    }
    return root
}