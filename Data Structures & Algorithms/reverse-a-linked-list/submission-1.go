/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	next := head.Next
	if next == nil {
		return head
	}
	link := next.Next
	head.Next = nil
	for {
		next.Next = head
		head = next
		next = link
		if link == nil {
			return head
		}
		link = next.Next
	}
}
