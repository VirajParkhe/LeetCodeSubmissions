/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type node struct {
    val int
    level int
}

func trav(val []node, l,level int, ind *int) *TreeNode {
     var root *TreeNode
    if *ind >= l {
        return root
    }
    //fmt.Println(*ind, level)
    if level == val[*ind].level {
        root = &TreeNode{val[*ind].val,nil,nil}
        (*ind)++
        root.Left = trav(val,l,level+1, ind)
        root.Right = trav(val,l,level+1, ind)
    }
    return root
}

func recoverFromPreorder(traversal string) *TreeNode {
    values := []node{}
    start :=0
    found := true
    level := 0
    for i:=0;i<len(traversal);i++{
        if (traversal[i] == '-') && (!found) {
            val, _ := strconv.Atoi(traversal[start:i])
            values = append(values, node{val, level})
            level = 1
            found  = true   
        }else if (traversal[i] != '-') && found{
            start = i
            found = false
        }else if traversal[i] == '-' {
            level++
        }
    }
    val, _ := strconv.Atoi(traversal[start:])
    values = append(values, node{val, level})
    //fmt.Println(values)
    ind := 0
    return trav(values, len(values), 0, &ind)
    
}