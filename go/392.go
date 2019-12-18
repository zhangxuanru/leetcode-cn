/*
@Time : 2019/12/16 16:36
@Author : zxr
@File : 392
@Software: GoLand
*/
package _go

import "strings"

/**
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
示例 1:
s = "abc", t = "ahbgdc"
返回 true.

示例 2:
s = "axc", t = "ahbgdc"
返回 false.
后续挑战 :

如果有大量输入的 S，称作S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/is-subsequence
 **/

//todo： 明天重写
func IsSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	letterMap := make(map[rune]struct{})
	upperMap := make(map[rune]struct{})
	for i := 'a'; i <= 'z'; i++ {
		letterMap[i] = struct{}{}
	}
	for i := 'A'; i <= 'Z'; i++ {
		upperMap[i] = struct{}{}
	}
	for _, v := range s {
		r := rune(v)
		delete(letterMap, r)
		delete(upperMap, r)
	}
	if len(letterMap) != 26 {
		for v := range letterMap {
			t = strings.ReplaceAll(t, string(v), "")
		}
	}
	if len(upperMap) != 26 {
		for v := range upperMap {
			t = strings.ReplaceAll(t, string(v), "")
		}
	}
	if len(s) == len(t) {
		return s == t
	}
	i := 0
	j := 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

//todo 推荐 这种
func IsSubsequence2(s string, t string) bool {
	i := 0
	j := 0
	sLen := len(s)
	tLen := len(t)
	for i < sLen && j < tLen {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == sLen
}
