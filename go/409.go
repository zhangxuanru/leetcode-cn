/*
@Time : 2019/12/3 11:55
@Author : zxr
@File : 409
@Software: GoLand
*/
package _go

/**
给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
注意:
假设字符串的长度不会超过 1010。

示例 1:

输入:
"abccccdd"

输出:
7

解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindrome
**/

func LongestPalindrome(s string) int {
	i := 0
	sLen := len(s)
	maxLen := 0
	isSingle := 0
	strMap := make(map[byte]int)
	for i < sLen {
		strMap[s[i]]++
		i++
	}
	//if len(strMap) == 1 {
	//	return sLen
	//}
	for key, count := range strMap {
		if count%2 == 0 {
			maxLen += count
		} else if count > 2 {
			maxLen += count - 1
			isSingle = 1
		} else {
			isSingle = 1
		}
		delete(strMap, key)
	}
	return maxLen + isSingle
}
