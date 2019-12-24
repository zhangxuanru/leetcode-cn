/*
@Time : 2019/12/23 19:07
@Author : zxr
@File : 125
@Software: GoLand
*/
package _go

import (
	"bytes"
	"strings"
)

/**
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。
示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-palindrome
 **/
func IsPalindromeFive(s string) bool {
	s = strings.ToLower(s)
	var s1Buf bytes.Buffer
	var s2Buf bytes.Buffer
	for _, r := range s {
		if (r >= 48 && r <= 57) || (r >= 'a' && r <= 'z') {
			s1Buf.WriteRune(r)
		}
	}
	j := len(s) - 1
	for j >= 0 {
		if (s[j] >= 48 && s[j] <= 57) || (s[j] >= 'a' && s[j] <= 'z') {
			s2Buf.WriteByte(s[j])
		}
		j--
	}
	return s1Buf.String() == s2Buf.String()
}
