package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// This solution has less mechanical sympathy (way more cache misses), but doesn't need any additional memory
	count := n
	previousToNth := head
	for cursor := head; cursor != nil; cursor = cursor.Next {
		if count < 0 {
			previousToNth = previousToNth.Next
		}
		count--
	}
	if count == 0 {
		return head.Next
	} else {
		previousToNth.Next = previousToNth.Next.Next
		return head
	}
}
