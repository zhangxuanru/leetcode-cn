/*
@Time : 2019/9/23 14:41
@Author : zxr
@File : mergeTwoLists
@Software: GoLand
*/
package lettcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newNode := new(ListNode)
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
	return newNode.Next
}
