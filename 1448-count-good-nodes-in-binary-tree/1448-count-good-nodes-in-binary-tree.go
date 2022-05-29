/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    ret := 0
    max := root.Val
    findGoodNodes(root, nil, &max,&ret)
    return ret
}

func findGoodNodes(root, parent *TreeNode, max, count *int) {
    if root == nil {
        return
    }
    if parent == nil || ( (root.Val >= *max) ) {
        *count +=1
    }
    if root.Val > *max {
        *max = root.Val
    }
    temp := *max
    findGoodNodes(root.Left, root, max, count)
    *max = temp
    findGoodNodes(root.Right, root, max, count)
    *max=temp
}