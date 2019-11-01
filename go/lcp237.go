/*
@Time : 2019/10/30 14:31
@Author : zxr
@File : lcp237
@Software: GoLand
*/
package _go

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteNode(node *ListNode) {
	next := node.Next
	for next.Next != nil {
		node.Val = next.Val
		node = next
		next = node.Next
	}
	node.Val = next.Val
	node.Next = nil
	//intList := []int{9, 1, 5, 4}
	//head := &ListNode{}
	//for _, val := range intList {
	//	newNode := &ListNode{Val: val}
	//	newNode.Next = head
	//	head = newNode
	//}

}
