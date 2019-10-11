package main

import "fmt"

// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

// 示例：

// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) Print() {

	for {
		fmt.Println("v = ", ln.Val)
		if ln.Next == nil {
			break
		}
		ln = ln.Next

	}
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	rListNode := &ListNode{}
	tempNode := rListNode
	flagAdd := false
	for {
		rs := l1.Val + l2.Val
		if flagAdd {
			rs++
			flagAdd = false
		}
		fmt.Println("rs=", rs)
		if rs < 10 {
			tempNode.Val = rs
		} else {
			tempNode.Val = rs % 10
			flagAdd = true
		}
		l1 = l1.Next
		l2 = l2.Next
		if l1 == nil || l2 == nil {
			break
		}
		tempNode.Next = &ListNode{}
		tempNode = tempNode.Next
	}
	return rListNode

}
func main() {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}

	sum := addTwoNumbers(l1, l2)
	sum.Print()

}
