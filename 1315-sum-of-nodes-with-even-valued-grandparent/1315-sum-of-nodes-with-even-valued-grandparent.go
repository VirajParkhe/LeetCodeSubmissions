/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var total int
func sum(root, parent, gparent *TreeNode) {
    if root == nil {
        return
    }
    if gparent != nil && gparent.Val % 2 == 0 {
        total += root.Val
    }
    sum(root.Left, root, parent)
    sum(root.Right, root, parent)
}

func sumEvenGrandparent(root *TreeNode) int {
    total = 0
    sum(root,nil, nil)
    return total
}