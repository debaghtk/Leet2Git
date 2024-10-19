# Find Center of Star Graph
# Difficulty: Easy
# Language: golang
# Topic: Graph
# Tags: Graph
# Link: https://leetcode.com/problems/find-center-of-star-graph/

func findCenter(edges [][]int) int {
    edge1 := edges[0]
    edge2 := edges[1]
    
    if edge1[0] == edge2[0] ||  edge1[0] == edge2[1]{
        return edge1[0]
    }
    return edge1[1]
}