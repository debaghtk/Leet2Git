# Find the Town Judge
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/find-the-town-judge/

func findJudge(n int, trust [][]int) int {
    // town judge has no neighbours
    // town judge is a neighbour to everyone

    graph := map[int][]int{}
    reverseGraph := map[int][]int{}
    for _,v := range trust {
        graph[v[0]] = append(graph[v[0]], v[1])
        reverseGraph[v[1]] = append(reverseGraph[v[1]], v[0])
    }

    for i:=1;i<=n;i++ {
        if len(graph[i])==0 && len(reverseGraph[i])==n-1{
            return i
        }
    }
    return -1
}