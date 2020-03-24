/*
@Time : 2020/3/24 12:02
@Author : zxr
@File : 25
@Software: GoLand
*/
package ListNode

import "muke/lettcode/go/offer/config"

/**
链表合表
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
限制：

0 <= 链表长度 <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof
*/
func MergeTwoLists(l1 *config.ListNode, l2 *config.ListNode) *config.ListNode {
	list := new(config.ListNode)
	cur := list
	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 == nil {
			cur.Next = l2
			break
		}
		if l2 == nil {
			cur.Next = l1
			break
		}
		if l1.Val > l2.Val {
			cur.Next = l2
			l2 = l2.Next
			cur = cur.Next
		} else {
			cur.Next = l1
			l1 = l1.Next
			cur = cur.Next
		}
	}
	return list.Next
}
