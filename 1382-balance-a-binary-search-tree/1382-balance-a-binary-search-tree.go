/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var nodes []int
func traverse(root *TreeNode) {
    if root == nil {
        return
    }    
    traverse(root.Left)
    nodes = append(nodes, root.Val)
    traverse(root.Right)
}

func build(nodes []int) *TreeNode{
    l := len(nodes)
    if l == 0 {
        return nil
    }    
    var lt, rt *TreeNode
    mid := nodes[l/2]
    if l/2 - 1 >=0 {
        lt = build(nodes[:l/2])
    }
    if l/2 != l -1 {
        rt = build(nodes[l/2 +1 :])
    }
    return &TreeNode{mid,lt,rt}
}

func balanceBST(root *TreeNode) *TreeNode {
    nodes =  []int{}
    traverse(root)
    return build(nodes)
}