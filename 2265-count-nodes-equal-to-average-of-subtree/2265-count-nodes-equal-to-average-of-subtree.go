/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var total int
func calcAvg(root *TreeNode)( int, int ) {
    if root==nil {
        return 0, 0
    }
    l, ln := calcAvg(root.Left)
    r, rn := calcAvg(root.Right)
    
    if (l+r+root.Val)/(ln+rn+1) == root.Val {
        total++
    }
    return (root.Val + l + r), 1+ln+rn
    
}
func averageOfSubtree(root *TreeNode) int {
    total = 0
    _, _ = calcAvg(root)
    return total
}