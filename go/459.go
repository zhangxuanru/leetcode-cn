/*
@Time : 2019/12/18 20:52
@Author : zxr
@File : 459
@Software: GoLand
*/
package _go

import (
	"fmt"
	"strings"
)

/**
给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
示例 1:
输入: "abab"
输出: True

解释: 可由子字符串 "ab" 重复两次构成。
示例 2:
输入: "aba"
输出: False
示例 3:

输入: "abcabcabcabc"
输出: True
解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/repeated-substring-pattern
**/

//func 2
func RepeatedSubstringPattern2(s string) bool {
	fmt.Println(s)
	sLen := len(s)
	left := 0
	right := 0
	var tmpBty string
	for left < sLen {
		if right > sLen {
			right = sLen - 1
		}
		fmt.Println("--eq:", tmpBty, "left----", left, "right---:", right, "rstr:", string(s[left:right]))
		if len(tmpBty) > 0 && tmpBty == s[left:right] {
			num := sLen / len(tmpBty)
			if strings.EqualFold(strings.Repeat(tmpBty, num), s) {
				return true
			}
		}
		tmpBty += string(s[left])
		right = left + len(tmpBty) + 1
		left++
	}
	return false
}

//方法2
func RepeatedSubstringPattern(s string) bool {
	ns := s + s
	return strings.Contains(ns[1:len(ns)-1], s)
}

//方法3
func repeatedSubstringPattern(str string) bool {
	s := str + str
	s = s[1 : len(s)-1]
	return strings.Contains(s, str)
}

//
////方法4
//func RepeatedSubstringPatternTest(str string) bool {
//	fmt.Println(str)
//	z := `/((\w).*?)/`
//	//`^(\w+)\1+$`
//	compile := regexp.MustCompile(z)
//	s := compile.FindAllStringSubmatch(str, -1)
//	findString := compile.Find([]byte(str))
//	allString := compile.FindAllString(str, -1)
//	fmt.Println("----", s, "---", findString, "--", allString)
//	return false
//}
