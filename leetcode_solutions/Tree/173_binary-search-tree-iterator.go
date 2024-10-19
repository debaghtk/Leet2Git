# Binary Search Tree Iterator
# Difficulty: Medium
# Language: golang
# Topic: Tree
# Tags: Stack, Tree, Design, Binary Search Tree, Binary Tree, Iterator
# Link: https://leetcode.com/problems/binary-search-tree-iterator/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    current int
    traversal []int
}

func Constructor(root *TreeNode) BSTIterator {
    return BSTIterator{
        current: 0,
        traversal: inorder(root),
    }
}


func (this *BSTIterator) Next() int {
    //this.current++
    answer := this.traversal[this.current]
    this.current++
    return answer
}


func (this *BSTIterator) HasNext() bool {
    if this.current >= len(this.traversal) {
        return false
    }
    return true
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */


func inorder(root *TreeNode) []int {
    answer := []int{}
    if root == nil{
        return answer
    }
    if root.Left!=nil{
        answer = append(answer, inorder(root.Left)...)
    }
    
    answer = append(answer, root.Val)

    if root.Right != nil {
        answer = append(answer, inorder(root.Right)...)
    }
    
    return answer
}