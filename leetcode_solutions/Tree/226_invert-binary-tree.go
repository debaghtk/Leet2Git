# Invert Binary Tree
# Difficulty: Easy
# Language: golang
# Topic: Tree
# Tags: Tree, Depth-First Search, Breadth-First Search, Binary Tree
# Link: https://leetcode.com/problems/invert-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil || (root.Left == nil && root.Right == nil) {
        return root
    }
    // left := root.Left
   // right := root.Right

    left := invertTree(root.Right)
    right := invertTree(root.Left)

    root.Left = left
    root.Right = right
    return root    
}