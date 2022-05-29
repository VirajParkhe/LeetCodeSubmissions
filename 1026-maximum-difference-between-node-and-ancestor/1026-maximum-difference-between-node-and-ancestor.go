/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "math"
var max int
func maxAncestorDiff(root *TreeNode) int {
     max = -1
    _, _ = findMinMax(root)
    return max
}


func findMinMax(root *TreeNode) (int, int) {
    if root == nil {
        return 100001, -1
    }
    if root.Left ==  nil && root.Right == nil{
        return root.Val, root.Val
    }
    lm, lmx := findMinMax(root.Left)
    rm, rmx := findMinMax(root.Right)
    low, high := lm, lmx
    if rm < lm {
        low=rm
    }
    
    if rmx > lmx {
        high = rmx
    }
    if int(math.Abs(float64(root.Val)- float64(low))) > max {
        fmt.Println(root.Val, low, high)
        max = int(math.Abs(float64(root.Val)- float64(low)))
    }
    if int(math.Abs(float64(root.Val)- float64(high))) > max {
        fmt.Println(root.Val, low, high)
        max = int(math.Abs(float64(root.Val)- float64(high)))
    }
    
    if root.Val < low {
        low = root.Val    
    }
    if root.Val > high {
        high= root.Val
    }
    return low, high
}