package main

import "fmt"

// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

//

// 示例:

// 给定 1->2->3->4, 你应该返回 2->1->4->3.

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	tmp := head
	for {
		if tmp == nil || tmp.Next == nil {
			break
		}
		tmp.Val, tmp.Next.Val = tmp.Next.Val, tmp.Val
		tmp = tmp.Next.Next
	}
	return head
}

func (ln *ListNode) printListNode() {
	for {
		fmt.Println(" value :", ln.Val)
		if ln.Next == nil {
			break
		}
		ln = ln.Next
	}
}

func main() {
	ln := &ListNode{Val: 1}
	ln.Next = &ListNode{Val: 2}
	ln.Next.Next = &ListNode{Val: 3}
	ln.Next.Next.Next = &ListNode{Val: 4}
	ln.printListNode()
	fmt.Println()
	tln := swapPairs(ln)
	tln.printListNode()

}
