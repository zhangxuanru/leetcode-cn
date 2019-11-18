/*
@Time : 2019/11/18 17:50
@Author : zxr
@File : 206
@Software: GoLand
*/
package _go

import "fmt"

/**
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func ReverseList(head *ListNode) *ListNode {
	tList := &ListNode{}
	if head == nil || head.Next == nil {
		return head
	}
	setTail := false
	for head != nil {
		if setTail == false {
			tList.Val = head.Val
			tList.Next = nil
			setTail = true
		} else {
			tmp := &ListNode{}
			tmp.Val = head.Val
			tmp.Next = tList
			tList = tmp
		}
		head = head.Next
	}

	for tList != nil {
		fmt.Println("Val:", tList.Val)
		tList = tList.Next
	}
	return tList
}
