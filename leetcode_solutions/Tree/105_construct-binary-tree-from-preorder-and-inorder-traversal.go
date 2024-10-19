# Construct Binary Tree from Preorder and Inorder Traversal
# Difficulty: Medium
# Language: golang
# Topic: Tree
# Tags: Array, Hash Table, Divide and Conquer, Tree, Binary Tree
# Link: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    answer := &TreeNode{}
    if len(preorder) == 0 {
        return nil
    }
    root := preorder[0]
    answer.Val = root
    if len(preorder) == 1 {
        return answer
    }
    // find index of root in inorder
    index:=0
    for i,v := range inorder {
        if v == root {
            index = i
            break
        }
    }
    //inorderRight := inorder[index+1:]
    inorderLeft := inorder[:index]
    sizeLeft := len(inorderLeft)
   // sizeRight := len(inorderRight)
    answer.Right = buildTree(preorder[1+sizeLeft:], inorder[index+1:])
    answer.Left = buildTree(preorder[1:1+sizeLeft], inorder[:index])
    return answer
}