# Course Schedule II
# Difficulty: Medium
# Language: golang
# Topic: Graph
# Tags: Depth-First Search, Breadth-First Search, Graph, Topological Sort
# Link: https://leetcode.com/problems/course-schedule-ii/

func findOrder(numCourses int, prerequisites [][]int) []int {
    graph := map[int][]int{}
    //indegree := make([]int, numCourses)
    answer := []int{}

    for _,v := range prerequisites {
        graph[v[1]] = append(graph[v[1]], v[0])
        //indegree[v[0]]++
    }

    visited := make([]int, numCourses)

    var sort func(node int) bool
    sort = func(node int) bool {
        if visited[node]==1{
            return false
        }
        if visited[node]==2{
            return true
        }
        
        visited[node]=1
        for _, neighbour := range graph[node]{
           if !sort(neighbour) {
            return false
           }
        }

        visited[node]=2
        answer = append([]int{node}, answer...)
        return true
    }

    for i,v := range visited {
        if v == 0 {
            if !sort(i){
               return []int{} 
            }
        }
    }

    return answer
}