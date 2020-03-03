/*
@Time : 2020/3/3 18:59
@Author : zxr
@File : 58
@Software: GoLand
*/
package go_interview

import (
	"bytes"
)

/**
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，
该函数将返回左旋转两位得到的结果"cdefgab"。

示例 1：
输入: s = "abcdefg", k = 2
输出: "cdefgab"
示例 2：

输入: s = "lrloseumgh", k = 6
输出: "umghlrlose"

限制：
1 <= k < s.length <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof
*/
func ReverseLeftWords(s string, n int) string {
	var buf bytes.Buffer
	buf.WriteString(s[n:])
	buf.WriteString(s[0:n])
	return buf.String()
}
