/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var path []int
func trav(root *TreeNode){
    if root != nil {
        path = append(path, root.Val)
        trav(root.Left )
       trav(root.Right )
    }
}

func preorderTraversal(root *TreeNode) []int {
    path = []int{}
    trav(root)
    return path
}