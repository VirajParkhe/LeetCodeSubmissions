/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
    end := -1
    l := len(preorder)
    if l == 0 {
        return nil
    }
    for i:=1;i<l;i++{
        if preorder[i] > preorder[0]{
            end = i
            break
        }
    }
    var lt, rt *TreeNode
    if end != -1 {
        rt = bstFromPreorder(preorder[end:])
    }
    if l > 1 {
        if end == -1 {
            lt = bstFromPreorder(preorder[1:])
        }else{
             lt = bstFromPreorder(preorder[1:end])
        }
    }
    return &TreeNode{preorder[0],lt,rt}
}
