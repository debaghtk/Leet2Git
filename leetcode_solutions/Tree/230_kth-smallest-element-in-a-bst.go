# Kth Smallest Element in a BST
# Difficulty: Medium
# Language: golang
# Topic: Tree
# Tags: Tree, Depth-First Search, Binary Search Tree, Binary Tree
# Link: https://leetcode.com/problems/kth-smallest-element-in-a-bst/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    sorted := []int{}
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode){  
        if node == nil || len(sorted) == k {
            return
        }
        dfs(node.Left)
        sorted = append(sorted, node.Val)
        dfs(node.Right)
    }

    dfs(root)
    return sorted[k-1]
}