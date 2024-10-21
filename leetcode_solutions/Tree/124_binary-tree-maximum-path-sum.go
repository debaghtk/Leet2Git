# Binary Tree Maximum Path Sum
# Difficulty: Hard
# Language: golang
# Topic: Tree
# Tags: Dynamic Programming, Tree, Depth-First Search, Binary Tree
# Link: https://leetcode.com/problems/binary-tree-maximum-path-sum/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    answer := -1001
    var maxSumWithNode func(node *TreeNode) int
    maxSumWithNode = func(node *TreeNode) int{
        if node == nil{
            return 0
        }
        left := max(maxSumWithNode(node.Left),0)
        right := max(maxSumWithNode(node.Right),0)
        currentMax := node.Val+left+right
        if currentMax > answer{
            answer = currentMax
        }
        return node.Val + max(left, right)
    }

    maxSumWithNode(root)
    return answer
}

func max (a,b int) int {
    if a>b {
        return a
    }
    return b
}