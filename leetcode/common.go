package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func toBinaryTree(s string) *TreeNode {
	// See https://support.leetcode.com/hc/en-us/articles/360011883654-What-does-1-null-2-3-mean-in-binary-tree-representation-
	trimmed := strings.Trim(s, "[]{}")
	if len(trimmed) == 0 {
		return nil
	}

	var nodes []*TreeNode
	for _, split := range strings.Split(trimmed, ",") {
		var node *TreeNode
		x := strings.TrimSpace(split)
		if x != "null" {
			val, err := strconv.Atoi(x)
			if err != nil {
				fmt.Println("Error parsing integer: " + x)
			}
			node = &TreeNode{Val: val}
		}
		nodes = append(nodes, node)
	}

	root := nodes[0]
	kid := 1
	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		if node != nil {
			if kid < len(nodes) {
				node.Left = nodes[kid]
				kid++
				if kid < len(nodes) {
					node.Right = nodes[kid]
					kid++
				}
			}
		}
	}
	return root
}

func fromBinaryTree(root *TreeNode) string {
	var elems []string

	level := []*TreeNode{root}
	for len(level) != 0 {
		var next []*TreeNode
		nulls := 0
		for _, node := range level {
			if node == nil {
				nulls++
			} else {
				for i := 0; i < nulls; i++ {
					elems = append(elems, "null")
				}
				nulls = 0
				elems = append(elems, strconv.Itoa(node.Val))
				next = append(next, node.Left, node.Right)
			}
		}
		level = next
	}

	return "[" + strings.Join(elems, ",") + "]"
}
