# Linked List Cycle
# Difficulty: Easy
# Language: golang
# Topic: Linked List
# Tags: Hash Table, Linked List, Two Pointers
# Link: https://leetcode.com/problems/linked-list-cycle/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil{
        return false
    }
    
    i,j:=head,head
    for {
        if i.Next==nil {
            return false
        }
        i=i.Next
        if j.Next ==nil || j.Next.Next == nil{
            return false
        }
        j=j.Next.Next
        
        if i==j{
            return true
        }
    }
    
    return false
}

// solution 1 - create array of addresses and check for duplicates