/*
@Time : 2020/3/24 11:16
@Author : zxr
@File : 24
@Software: GoLand
*/
package ListNode

import (
	"muke/lettcode/go/offer/config"
)

/**
链表反转
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL


限制：

0 <= 节点个数 <= 5000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof
*/

func ReverseList(head *config.ListNode) *config.ListNode {
	var pre *config.ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

// func2  链表反转
func ReverseList2(head *config.ListNode) *config.ListNode {
	pre := &config.ListNode{}
	tail := false
	for head != nil {
		if tail == false {
			pre.Val = head.Val
			pre.Next = nil
			tail = true
		} else {
			tmp := &config.ListNode{}
			tmp.Val = head.Val
			tmp.Next = pre
			pre = tmp
		}
		head = head.Next
	}
	return pre
}
