# Lowest Common Ancestor of a Binary Tree
# Difficulty: Medium
# Language: golang
# Link: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  //lca := root
  var left, right *TreeNode
  if nodeispq(root,p,q) {
    return root
  }
  if root.Left == nil && root.Right == nil{
    return nil
  }
  if root.Left != nil {
    left = lowestCommonAncestor(root.Left, p, q)
  } 
  if root.Right != nil {
    right = lowestCommonAncestor(root.Right, p, q)
  }

  if left == nil && right != nil {
    return right
  }
  if left != nil && right == nil {
    return left
  }
  if left == nil && right == nil {
    return nil
  }
  return root
}

func nodeispq(node, p,q *TreeNode) bool {
    if node.Val == p.Val || node.Val == q.Val {
        return true
    }
    return false
}