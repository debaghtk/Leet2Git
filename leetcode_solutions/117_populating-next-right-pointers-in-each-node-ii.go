# Populating Next Right Pointers in Each Node II
# Difficulty: Medium
# Language: golang
# Link: https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/

func connect(root *Node) *Node {
    //return root
    if root == nil || (root.Left == nil && root.Right == nil){
        return root
    }
    // depth = 2
    if root.Left !=nil {
        root.Left.Next = assignNextForLeftChild(root)
    }
    if root.Right !=nil {
        root.Right.Next = assignNextForRightChild(root)
    }

    root.Right = connect(root.Right)
    root.Left = connect(root.Left)
    return root
}

func assignNextForRightChild(root *Node) *Node {
    if root.Next == nil {
        return nil
    }
    for next := root.Next;next!=nil;next=next.Next{
        if next.Left != nil {
            return next.Left
        }
        if next.Right != nil{
            return next.Right
        }
    }
    return nil
}

func assignNextForLeftChild(root *Node) *Node{
    if root.Right != nil{
        return root.Right
    }
    return assignNextForRightChild(root)
}