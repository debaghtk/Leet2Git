# Maximum Depth of Binary Tree
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/maximum-depth-of-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 //import "math"
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1+max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a,b int) int{
 if a > b {
    return a
 }
 return b
}