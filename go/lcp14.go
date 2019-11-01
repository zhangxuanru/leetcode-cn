/*
@Time : 2019/9/20 14:59
@Author : zxr
@File : longestCommonPrefix
@Software: GoLand
*/
/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
*/
package _go

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for _, str := range strs {
			if len(str) <= i || str[i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
