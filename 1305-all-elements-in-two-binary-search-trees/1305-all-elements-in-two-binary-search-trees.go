/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var ele1, ele2 []int
func traverse(root *TreeNode, isFirst bool) {
    if root==nil{
        return
    }
    traverse(root.Left, isFirst)
    if isFirst{
        ele1 = append(ele1, root.Val)   
    }else{
        ele2 = append(ele2, root.Val)
    }
    traverse(root.Right, isFirst)
    return
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    ele1 = []int{}
    ele2 = []int{}
    ret := []int{}
    traverse(root1, true)
    traverse(root2, false)
    var i, j, l1, l2 int
    l1 = len(ele1)
    l2 = len(ele2)
    i,j = 0,0
    for ;i<l1;{
        if j<l2 {
            if ele1[i] < ele2[j] {
                ret = append(ret, ele1[i])
                i++
            }else{
                ret = append(ret, ele2[j])
                j++
            }
        }else{
            ret = append(ret, ele1[i])
            i++
        }
    }
    if j < l2 {
        for ;j<l2;j++{
            ret = append(ret, ele2[j])
        }
    }
    return ret
}