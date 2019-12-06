/*
@Time : 2019/12/6 15:48
@Author : zxr
@File : 796
@Software: GoLand
*/
package _go

/**
给定两个字符串, A 和 B。

A 的旋转操作就是将 A 最左边的字符移动到最右边。
例如, 若 A = 'abcde'，在移动一次之后结果就是'bcdea' 。
如果在若干次旋转操作之后，A 能变成B，那么返回True。

示例 1:
输入: A = 'abcde', B = 'cdeab'
输出: true

示例 2:
输入: A = 'abcde', B = 'abced'
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-string
**/

func RotateString(A string, B string) bool {
	aLen := len(A)
	i := 0
	if A == B {
		return true
	}
	for i < aLen {
		A += string(A[0])
		A = A[1:]
		if A == B {
			return true
		}
		i++
	}
	return false
}
