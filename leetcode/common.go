package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func toList(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	head := &ListNode{Val: list[0]}
	cursor := head
	for i := 1; i < len(list); i++ {
		cursor.Next = &ListNode{Val: list[i]}
		cursor = cursor.Next
	}
	return head
}

func fromList(head *ListNode) []int {
	list := make([]int, 0, 30)
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	return list
}

type Void struct{}

var void Void
