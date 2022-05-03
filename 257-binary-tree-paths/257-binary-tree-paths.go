/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    if root == nil {
        return []string{}
    }
    paths := binaryTreePaths(root.Left)
    paths = append(paths, binaryTreePaths(root.Right)...)
    if len(paths) == 0 {
        return []string{strconv.Itoa(root.Val)}
    }
    for i:=0;i<len(paths);i++{
        paths[i] = strconv.Itoa(root.Val) + "->" + paths[i]
    }
    return paths
}