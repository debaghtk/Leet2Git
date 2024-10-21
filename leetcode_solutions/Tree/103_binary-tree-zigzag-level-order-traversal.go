# Binary Tree Zigzag Level Order Traversal
# Difficulty: Medium
# Language: golang
# Topic: Tree
# Tags: Tree, Breadth-First Search, Binary Tree
# Link: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

func zigzagLevelOrder(root *TreeNode) [][]int {
    answer := [][]int{}
    stack := []*TreeNode{}
    if root!=nil{
        stack = append(stack, root)
    }
    var bfs func(stack []*TreeNode)

    bfs = func(stack []*TreeNode){
        length := len(stack)
        if length==0{
            return
        }
        level := []int{}
        order := len(answer)%2
        for i:=0;i<length;i++{
            if order == 0 {
                level = append(level, stack[i].Val)
            } else {
                level = append([]int{stack[i].Val}, level...)
            } 
            
            if stack[i].Left != nil{
                stack = append(stack, stack[i].Left)
            }
            if stack[i].Right != nil{
                stack = append(stack, stack[i].Right)
            }
        }
        answer = append(answer, level)
        stack=stack[length:]
        bfs(stack)
        return
    }

    bfs(stack)
    return answer
}