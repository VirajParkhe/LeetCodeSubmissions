/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trav(root *TreeNode, path []int) []int{
    if root != nil {
        path = append(path, root.Val)
        path= trav(root.Left, path)
        path =trav(root.Right, path)
    }
    return path
}

func preorderTraversal(root *TreeNode) []int {
    return trav(root, []int{})
}