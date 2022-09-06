package leetcode

import "container/heap"

func mergeKLists(lists []*ListNode) *ListNode {
	h := NewListNodeHeap(len(lists))
	heap.Init(h)
	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}

	root := &ListNode{}
	tail := root
	for h.Len() > 0 {
		next := heap.Pop(h).(*ListNode)
		if next.Next != nil {
			heap.Push(h, next.Next)
		}
		tail.Next = next
		tail = next
	}

	return root.Next
}

type ListNodeHeap []*ListNode

func NewListNodeHeap(n int) *ListNodeHeap {
	h := make(ListNodeHeap, 0, n)
	return &h
}

func (h ListNodeHeap) Len() int           { return len(h) }
func (h ListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h ListNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	last := len(old) - 1
	x := old[last]
	*h = old[:last]
	return x
}
