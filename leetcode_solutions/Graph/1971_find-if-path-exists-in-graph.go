# Find if Path Exists in Graph
# Difficulty: Easy
# Language: golang
# Topic: Graph
# Tags: Depth-First Search, Breadth-First Search, Union Find, Graph
# Link: https://leetcode.com/problems/find-if-path-exists-in-graph/

func validPath(n int, edges [][]int, source int, destination int) bool {
    if source == destination {
        return true
    }
    nn := map[int][]int{}
    for _,v := range edges {
        // if v[1]!=source{
            nn[v[0]]= append(nn[v[0]], v[1])
        //}
      //  if v[0]!=source{
            nn[v[1]]= append(nn[v[1]], v[0])
        //}
    }

    if len(nn[source])==0 || len(nn[destination])==0 {
        return false
    }
    
    visited:=map[int]bool{}
    for queue:=nn[source];len(queue)>0;queue = queue[1:]{
        node := queue[0]
        
        if visited[node] || len(nn[node])==0{
            continue
        }
        if node == destination || nodeInArray(destination, nn[node]) {
            return true
        }
        visited[node]=true
        for _, neighbor := range nn[node] {
            if !visited[neighbor] {
                queue = append(queue, neighbor)
            }
        }
    }
    return false
}

func nodeInArray(destination int, nodes []int) bool {
    for _,v := range nodes {
        if v == destination{
            return true
        }
    }
    return false
}
