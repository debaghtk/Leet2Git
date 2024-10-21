# Average of Levels in Binary Tree
# Difficulty: Easy
# Language: golang
# Topic: Tree
# Tags: Tree, Depth-First Search, Breadth-First Search, Binary Tree
# Link: https://leetcode.com/problems/average-of-levels-in-binary-tree/

func averageOfLevels(root *TreeNode) []float64 {
    //eps := math.Pow(10,-5)
    answer := []float64{}
    q := []*TreeNode{}
    q = append(q, root)

    var bfs func(queue []*TreeNode)
    bfs = func(queue []*TreeNode) {
        length := len(queue)
        if length == 0{
            return
        }
        sum := 0
        for i:=0;i<length;i++{
             if queue[i] == nil {
                 continue
             }
            sum += queue[i].Val
            if queue[i].Left != nil{
                queue = append(queue, queue[i].Left)
            }
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            }    
        }

        avg := float64(sum)/float64(length)
        answer = append(answer, avg)
        if len(queue) > length {
            queue = queue[length:]
            bfs(queue)
        }
        return
    }
    bfs(q)
    return answer
}