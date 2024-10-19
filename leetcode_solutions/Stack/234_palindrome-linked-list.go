# Palindrome Linked List
# Difficulty: Easy
# Language: golang
# Topic: Stack
# Tags: Linked List, Two Pointers, Stack, Recursion
# Link: https://leetcode.com/problems/palindrome-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head==nil{
        return false
    }
    if head.Next==nil{
        return true
    }
    
    if head.Val == getLastVal(head){
        return isPalindrome(remove(head))
    }
    return false
}

func getLastVal(head *ListNode) int{
    ptr:=head
    if ptr.Next==nil{
        return ptr.Val
    }
    
    if ptr.Next.Next==nil{
        answer := ptr.Next.Val
        ptr.Next=nil
        return answer 
    }
    
    for{
        if ptr.Next.Next == nil{
            answer := ptr.Next.Val
        ptr.Next=nil
        return answer 
        }
        ptr = ptr.Next
    }
    return 0
}

func remove(head *ListNode) *ListNode{
    if head.Next==nil{
      return &ListNode{}    
    }
    return head.Next
    
}