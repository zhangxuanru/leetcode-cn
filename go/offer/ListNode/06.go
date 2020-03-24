/*
@Time : 2020/3/23 18:25
@Author : zxr
@File : 06
@Software: GoLand
*/
package ListNode

import (
	"fmt"
	"muke/lettcode/go/offer/config"
)

/**
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

示例 1：

输入：head = [1,3,2]
输出：[2,3,1]
*/
func ReversePrint(head *config.ListNode) []int {
	var res, data []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	l := len(res) - 1
	i := 0
	data = make([]int, len(res))
	for l >= 0 {
		data[i] = res[l]
		i++
		l--
	}
	return data
}

/**
链表反转
*/
func R2(head *config.ListNode) {
	var pre *config.ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next

		fmt.Println(pre)
	}
	for pre != nil {
		fmt.Println(pre.Val, "---")
		pre = pre.Next
	}
}
