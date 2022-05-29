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
    findGoodNodes(root,&[]*TreeNode{}, &ret)
    return ret
}

func findGoodNodes(root *TreeNode, nodes *[]*TreeNode, count *int) {
    if root == nil {
        return
    }
    isGood := true
    for i:=0;i<len(*nodes);i++{
        if (*nodes)[i].Val > root.Val{
            isGood = false
            break
        }
    }
    if isGood {
        *count +=1
    }
    *nodes = append(*nodes, root)
    findGoodNodes(root.Left, nodes, count)
    findGoodNodes(root.Right, nodes, count)
    *nodes =(*nodes)[:len(*nodes)-1]
    return
}