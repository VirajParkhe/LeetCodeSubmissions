/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lca(root, p,q *TreeNode) (bool, bool, *TreeNode) {
    if root == nil  || p == nil || q == nil{
        return false, false, nil
    }    
    if root.Val == p.Val {
        return true, false, p
    }
    if root.Val == q.Val {
        return false, true, q
    }
    l1, l2, tl := lca(root.Left, p,q)
    if l1 == true && l2 == true {
        return l1, l2, tl  
    }
    if l1 == true && root.Val == q.Val{
        return true, true, root
    }
    if l2 == true && root.Val == p.Val{
        return true, true, root
    }
    
    r1, r2, tr := lca(root.Right, p,q)
    if r1 == true && r2 == true {
        return true, true, tr
    }
    
    if r1 == true && root.Val == q.Val {
        return true, true, root
    }
    if r2 == true && root.Val == p.Val {
        return true, true, root
    }
    if l1 == true && r2 == true {
        return true, true, root
    }
    if l2 == true && r1 == true {
        return true, true, root
    }
    if l1 == true || l2 == true{
        return l1, l2, tl
    }
    if r1 == true || r2 == true {
        return r1, r2, tr
    }
    
    return false, false, nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    _,_, ret := lca(root, p, q)
    return ret
}