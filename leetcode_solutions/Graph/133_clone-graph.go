# Clone Graph
# Difficulty: Medium
# Language: golang
# Topic: Graph
# Tags: Hash Table, Depth-First Search, Breadth-First Search, Graph
# Link: https://leetcode.com/problems/clone-graph/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    visited := map[int]bool{}
    stack := make([]*Node,101)
    return cloneGraphHelper(node, visited, stack)
}

func cloneGraphHelper(node *Node, visited map[int]bool, stack []*Node) *Node {
    if node ==nil || visited[node.Val] {
        return node
    }

    copy := Node{
        Val: node.Val,
    }

    visited[node.Val]=true
    stack[node.Val]=&copy

    for _,childNode := range node.Neighbors {
        if visited[childNode.Val] {
            copy.Neighbors = append(copy.Neighbors, stack[childNode.Val])
            continue
        }
        copy.Neighbors = append(copy.Neighbors, cloneGraphHelper(childNode, visited, stack))
    }
    
    return &copy
}