# Minimum Absolute Difference in BST
# Difficulty: Easy
# Language: golang
# Topic: Tree
# Tags: Tree, Depth-First Search, Breadth-First Search, Binary Search Tree, Binary Tree
# Link: https://leetcode.com/problems/minimum-absolute-difference-in-bst/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    answer := 100000
    var calculate func(node *TreeNode)
    calculate = func(node *TreeNode) {
        if node == nil{
            return
        }
        leftDiff := -1
        rightDiff := -1
        leftNode := node.Left
        rightNode := node.Right
        calculate(leftNode)
        calculate(rightNode)
        for true {  
            if (leftNode == nil || leftNode.Right==nil) && (rightNode == nil || rightNode.Left == nil) {
                break
            }
            if leftNode != nil && leftNode.Right !=nil {
                leftNode = leftNode.Right
            }
            if rightNode != nil &&rightNode.Left !=nil {
                rightNode = rightNode.Left
            }
        }   
        if leftNode!=nil{
            leftDiff = node.Val - leftNode.Val
        }
        if rightNode != nil {
            rightDiff = rightNode.Val - node.Val
        }
        

        current := min(leftDiff, rightDiff)
        answer = min(answer, current)
            
        return 
    }

    calculate(root)
    return answer
}

func min (a,b int) int {
    if a == -1 && b == -1 {
        return -1
    }
    if a == -1 {
        return b
    }
    if b == -1 {
        return a
    }
    if a < b {
        return a
    }
    return b
}