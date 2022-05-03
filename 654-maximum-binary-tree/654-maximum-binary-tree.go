/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(nodes []int) *TreeNode {
    max := nodes[0]
    maxi := 0
    l := len(nodes)
    for i:=0;i<l;i++{
        if nodes[i] > max {
            max = nodes[i]
            maxi = i
        }
    }
    var lt, rt *TreeNode
    if maxi != 0 {
        lt = buildTree(nodes[:maxi])
    }
    if maxi != l -1{
        rt = buildTree(nodes[maxi+1:])
    }
    return &TreeNode{max, lt, rt}
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
    return buildTree(nums)
}