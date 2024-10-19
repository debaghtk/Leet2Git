# Sum Root to Leaf Numbers
# Difficulty: Medium
# Language: golang
# Topic: Tree
# Tags: Tree, Depth-First Search, Binary Tree
# Link: https://leetcode.com/problems/sum-root-to-leaf-numbers/

func sumNumbers(root *TreeNode) int {
    return helper(root, root.Val) 
}

func helper (root *TreeNode, upar int) int {
    answer := 0
    if root.Left == nil && root.Right == nil {
        return upar
    }
    if root.Left!=nil{
        answer += helper(root.Left, upar*10 + root.Left.Val)
    }
    if root.Right!=nil{
        answer += helper(root.Right, upar*10 + root.Right.Val)
    } 
    return answer
}