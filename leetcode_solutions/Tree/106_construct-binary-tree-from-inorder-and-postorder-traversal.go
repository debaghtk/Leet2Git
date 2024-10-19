# Construct Binary Tree from Inorder and Postorder Traversal
# Difficulty: Medium
# Language: golang
# Topic: Tree
# Tags: Array, Hash Table, Divide and Conquer, Tree, Binary Tree
# Link: https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 import "slices"
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) ==0 || len(postorder)==0{
        return nil
    }
    // if len(inorder) ==1 || len(postorder)==1{
    //     return &TreeNode{Val: inorder[0]}
    // }
    index := slices.Index(inorder, postorder[len(postorder)-1])
    inorderLeft := inorder[:index]
    inorderRight := inorder[index+1:]
    sizeLeft := len(inorderLeft)

    postorderLeft:= postorder[0:sizeLeft]
    postorderRight:= postorder[sizeLeft:len(postorder)-1]
    return &TreeNode{
        Val: postorder[len(postorder)-1],
        Left: buildTree(inorderLeft, postorderLeft),
        Right: buildTree(inorderRight, postorderRight),
    }
}