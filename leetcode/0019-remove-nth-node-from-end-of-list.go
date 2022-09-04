package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// The requirements are quite relaxed, otherwise we'd go for an array-based ring size of n (or ceiling power of two)
	const MaxSize = 30
	l := make([]*ListNode, 0, MaxSize)
	for cursor := head; cursor != nil; cursor = cursor.Next {
		l = append(l, cursor)
	}
	if n == len(l) {
		return head.Next
	} else {
		nth := len(l) - n
		l[nth-1].Next = l[nth].Next
		return head
	}
}
