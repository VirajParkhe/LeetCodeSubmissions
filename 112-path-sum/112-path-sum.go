/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sum(root *TreeNode, targetSum int)bool {
    if root == nil && targetSum == 0 {
        fmt.Println("a")
        return true
    }else if root == nil && targetSum != 0 {
        fmt.Println("b", targetSum)
        return false
    }
    if root.Left != nil {
        l := sum(root.Left, targetSum - root.Val)
        if l == true {
            return l
        }
    }
    if root.Right != nil {
        return sum(root.Right, targetSum - root.Val)
    }
    if root.Val == targetSum && root.Left == nil && root.Right == nil {
        return true
    }
    
    return false
}

func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    return sum(root, targetSum)
}