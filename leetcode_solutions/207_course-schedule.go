# Course Schedule
# Difficulty: Medium
# Language: golang
# Link: https://leetcode.com/problems/course-schedule/

func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make(map[int][]int)

    // Build the graph from prerequisites
    for _, v := range prerequisites {
        graph[v[1]] = append(graph[v[1]], v[0])
    }

    // State of each node: 0 = Not visited, 1 = Visiting, 2 = Visited
    visits := make([]int, numCourses)

    // DFS function
    var dfs func(node int) bool
    dfs = func(node int) bool {
        if visits[node] == 1 { // Cycle detected
            return false
        }
        if visits[node] == 2 { // Already processed
            return true
        }

        visits[node] = 1 // Mark as Visiting
        for _, neighbor := range graph[node] {
            if !dfs(neighbor) {
                return false
            }
        }
        visits[node] = 2 // Mark as Visited
        return true
    }

    // Perform DFS on each course
    for i := 0; i < numCourses; i++ {
        if visits[i] == 0 {
            if !dfs(i) {
                return false
            }
        }
    }

    return true
}
