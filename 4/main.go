package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	rListNode := &ListNode{}
	flagAdd := false
	for {
		rs := l1.Val + l2.Val
		if flagAdd {
			rs++
			flagAdd = false

		}
		if rs < 10 {
			rListNode.Val = rs
		} else {
			rListNodeVal = rs % 10
			flagAdd = true
		}
		rListNode.Next = &ListNode{}
		rListNode = rListNode.Next
		l1 = l1.Next
		l2 = l2.Next
		if l1.Next == nil || l2.Next == nil {
			break
		}
	}
	return rListNode

}
func main() {

}
