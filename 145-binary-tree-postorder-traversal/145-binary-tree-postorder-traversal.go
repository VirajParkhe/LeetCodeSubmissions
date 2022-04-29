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
        trav(root.Left)
        trav(root.Right)
        path = append(path, root.Val) 
    }
}

func postorderTraversal(root *TreeNode) []int {
    path = []int{}
    trav(root)
    return path
}