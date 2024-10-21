# Binary Tree Right Side View
# Difficulty: Medium
# Language: golang
# Topic: Tree
# Tags: Tree, Depth-First Search, Breadth-First Search, Binary Tree
# Link: https://leetcode.com/problems/binary-tree-right-side-view/

func rightSideView(root *TreeNode) []int {
    answer := []int{}
    if root == nil {
        return answer
    }
    queue := []*TreeNode{}
    queue = append(queue, root)

    var bfs func(q []*TreeNode) 
    bfs = func(q []*TreeNode) {
        if len(q) == 0 {
            return
        }

        answer = append(answer, q[len(q)-1].Val)
        length:= len(q)
        for i:=0; i<length;i++{
            if q[i].Left!=nil{
                q = append(q, q[i].Left)
            }
            if q[i].Right!=nil{
                q = append(q, q[i].Right)
            }
        }
        q = q[length:]
        bfs(q)
        return
    }

    bfs(queue)
    return answer
}