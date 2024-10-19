# Flatten Binary Tree to Linked List
# Difficulty: Medium
# Language: golang
# Link: https://leetcode.com/problems/flatten-binary-tree-to-linked-list/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    if root == nil || (root.Left==nil && root.Right==nil) {
        return
    }
    flatten(root.Left)
    flatten(root.Right)
    if last(root.Left)!=nil {
        last(root.Left).Right=root.Right
        root.Right=root.Left
    }
    root.Left=nil 
    return
}

func last(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    for answer:=root;;{
        if answer.Right == nil {
            return answer
        }
        answer = answer.Right
    }
    return nil
}