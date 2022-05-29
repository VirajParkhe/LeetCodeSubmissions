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
    findGoodNodes(root, nil, root.Val,&ret)
    return ret
}

func findGoodNodes(root, parent *TreeNode,max int, count *int) {
    if root == nil {
        return
    }
    if parent == nil || ( (root.Val >= max) ) {
        *count +=1
    }
    if root.Val > max {
        max = root.Val
    }
    findGoodNodes(root.Left, root, max, count)
    findGoodNodes(root.Right, root, max, count)
}