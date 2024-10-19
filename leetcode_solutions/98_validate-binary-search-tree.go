# Validate Binary Search Tree
# Difficulty: Medium
# Language: golang
# Link: https://leetcode.com/problems/validate-binary-search-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    if root == nil{
        return true
    }
    left := root.Left
    right := root.Right
    val := root.Val
    
    if root.Left == nil && root.Right == nil {
        return true
    }
    
    if root.Left == nil {
        validRightSide := validateRightSide(val, right)
        return validRightSide && isValidBST(right)
    }
    if root.Right == nil {
        validLeftSide := validateLeftSide(val, left)
        return validLeftSide && isValidBST(left)
    }
    
    validRightSide := validateRightSide(root.Val, root.Right)
    validLeftSide := validateLeftSide(root.Val, root.Left)
    
    return validLeftSide && validRightSide && isValidBST(root.Left) && isValidBST(root.Right)
}

func validateRightSide(val int, right *TreeNode) bool{
    if val >= right.Val {
        return false
    }
    if right.Left == nil && right.Right == nil{
        return true
    }
    if right.Left == nil{
        return validateRightSide(val, right.Right)
    }
    if right.Right == nil{
        return validateRightSide(val, right.Left)
    }
    return validateRightSide(val, right.Left) && validateRightSide(val, right.Right)
}

func validateLeftSide(val int,left *TreeNode) bool{
    if val <= left.Val {
        return false
    }
    if left.Left == nil && left.Right == nil{
        return true
    }
    if left.Left == nil{
        return validateLeftSide(val, left.Right)
    }
    if left.Right == nil{
        return validateLeftSide(val, left.Left)
    }
    return validateLeftSide(val,left.Left) && validateLeftSide(val,left.Right)
    
}