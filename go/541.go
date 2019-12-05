/*
@Time : 2019/12/5 20:54
@Author : zxr
@File : 541
@Software: GoLand
*/
package _go

/**
给定一个字符串和一个整数 k，你需要对从字符串开头算起的每个 2k 个字符的前k个字符进行反转。如果剩余少于 k 个字符，则将剩余的所有全部反转。如果有小于 2k 但大于或等于 k 个字符，则反转前 k 个字符，并将剩余的字符保持原样。

示例:

输入: s = "abcdefg", k = 2
输出: "bacdfeg"
要求:

该字符串只包含小写的英文字母。
给定字符串的长度和 k 在[1, 10000]范围内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-string-ii
 **/
func ReverseStr(s string, k int) string {
	sRune := []rune(s)
	sLen := len(s)
	for i := 0; i < sLen; i += 2 * k {
		if i+k < sLen {
			reverse(sRune, i, i+k-1)
		} else {
			reverse(sRune, i, sLen-1)
		}
	}
	return string(sRune)
}

func reverse(s []rune, start, end int) {
	for start < end {
		tmp := s[start]
		s[start] = s[end]
		s[end] = tmp
		start++
		end--
	}
}
