/*
@Time : 2019/12/26 10:54
@Author : zxr
@File : 21
@Software: GoLand
*/
package _go

/**
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
 **/
func MergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := &ListNode{}
	newNode := cur
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
			cur = cur.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
			cur = cur.Next
		}
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return newNode.Next
}

func MergeTwoLists3(l1 *ListNode, l2 *ListNode) *ListNode {
	newNode := &ListNode{}
	cur := newNode
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
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
			cur = cur.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
			cur = cur.Next
		}
	}
	return newNode.Next
}

//todo
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
