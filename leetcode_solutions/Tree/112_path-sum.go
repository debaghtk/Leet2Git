# Path Sum
# Difficulty: Easy
# Language: golang
# Topic: Tree
# Tags: Tree, Depth-First Search, Breadth-First Search, Binary Tree
# Link: https://leetcode.com/problems/path-sum/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    if root.Val == targetSum && root.Left == nil && root.Right == nil {
        return true
    }
    newTarget := targetSum - root.Val

    return hasPathSum(root.Left, newTarget) || hasPathSum(root.Right, newTarget)
}