/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traverse(root *TreeNode, val int) int {
    if root == nil {
        return -1
    } 
    r := traverse(root.Right, val)
    if r != -1 {
        root.Val += r
    }else {
        root.Val += val
    }
    l := traverse(root.Left, root.Val)
    if l == -1 {
        return root.Val
    }
    return l
}

func bstToGst(root *TreeNode) *TreeNode {
    _ = traverse(root, 0)
    return root
}