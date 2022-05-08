/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
    root *TreeNode
    exists map[int]bool
}

func construct(root *TreeNode, exists map[int]bool) {
    if root == nil {
        return
    }
    exists[root.Val] = true
    if root.Left != nil {
        root.Left.Val = 2*root.Val +1
    }
    if root.Right !=  nil {
        root.Right.Val = 2*root.Val +2
    }
    construct(root.Left, exists)
    construct(root.Right, exists)
}

func Constructor(root *TreeNode) FindElements {
    root.Val = 0
    exists := map[int]bool{}
    construct(root, exists)
    return FindElements{root, exists}
}


func (this *FindElements) Find(target int) bool {
    if _, ok := this.exists[target]; ok {
        return ok
    }
    return false
}


/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */