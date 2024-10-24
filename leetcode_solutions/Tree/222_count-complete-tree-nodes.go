# Count Complete Tree Nodes
# Difficulty: Easy
# Language: golang
# Topic: Tree
# Tags: Binary Search, Bit Manipulation, Tree, Binary Tree
# Link: https://leetcode.com/problems/count-complete-tree-nodes/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + countNodes(root.Left) + countNodes(root.Right)
}