# Reverse Linked List
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/reverse-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    
    if head==nil || head.Next==nil{
        return head
    }
    reversedLinkedList := reverseList(head.Next)
    AddAtTail(reversedLinkedList, head.Val)
    return reversedLinkedList
}

func AddAtTail(this *ListNode, val int)  {
	//if this.Val == 0 && this.Next == nil {
	//	this.Val = val
	//	return
	//}
	var ptr *ListNode
	for ptr = this; ptr.Next != nil; ptr = ptr.Next {
		continue
	}
	ptr.Next = &ListNode{Val: val}
}
