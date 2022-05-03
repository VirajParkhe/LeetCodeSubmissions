/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    children := []*TreeNode{root, nil}
    for ;; {
        added := false
        l := len(children)
        for _, val := range children {
            if val == nil {
                children = append(children, nil)
                break
            }
            if val.Left != nil {
                added = true
                children = append(children, val.Left)
            }
            if val.Right != nil {
                added = true
                children = append(children,val.Right )
            }
        }
         if !added{
            sum := 0
            for i:=len(children)-3;i>=0;i--{
        
                if children[i] == nil {
                    break
                }
                sum +=children[i].Val
            }
            return sum
        }
        children = children[l:]
       
       
        
    } 
    return -1
}