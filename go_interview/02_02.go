/*
@Time : 2020/3/3 16:28
@Author : zxr
@File : 02_02
@Software: GoLand
*/
package go_interview

/**
实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
注意：本题相对原题稍作改动

示例：

输入： 1->2->3->4->5 和 k = 2
输出： 4
说明：

给定的 k 保证是有效的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci
*/

//bad
func KthToLast(head *ListNode, k int) int {
	list := &ListNode{}
	tail := false
	retVal := 0

	for head != nil {
		if tail == false {
			list.Val = head.Val
			list.Next = nil
			tail = true
		} else {
			node := &ListNode{}
			node.Val = head.Val
			node.Next = list
			list = node
		}
		head = head.Next
	}
	for list != nil && k > 0 {
		retVal = list.Val
		list = list.Next
		k--
	}
	return retVal
}

//good   使用快慢指针
func KthToLast2(head *ListNode, k int) int {
	slow, fast := head, head
	for ; k > 0; k-- {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow.Val
}

//good
func KthToLast3(head *ListNode, k int) int {
	// 使用快慢指针
	p, q := head, head
	for q != nil {
		q = q.Next
		k--
		if k < 0 {
			p = p.Next
		}
	}
	return p.Val
}
