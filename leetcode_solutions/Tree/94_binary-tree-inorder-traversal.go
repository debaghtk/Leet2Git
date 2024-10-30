# Binary Tree Inorder Traversal
# Difficulty: Easy
# Language: golang
# Topic: Tree
# Tags: Stack, Tree, Depth-First Search, Binary Tree
# Link: https://leetcode.com/problems/binary-tree-inorder-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    answer:=[]int{}
    if root == nil{
        return answer
    }

    var pre *TreeNode
    node := root

    for true {
        if node == nil {
            break
        }
        if node.Left == nil {
            answer = append(answer, node.Val)
            node = node.Right
            continue
        }

        // find rightmost node of left subtree which is also known as predecessor
        pre = node.Left
        for true {
            if pre.Right == node { // second visit
                pre.Right = nil
                answer = append(answer, node.Val)
                node = node.Right
                break
            }
            if pre.Right != nil{
                pre = pre.Right
            }
            if pre.Right == nil{
                pre.Right = node
                node = node.Left
               // node = node.Left
                break
            }
        }
    }
    return answer
}