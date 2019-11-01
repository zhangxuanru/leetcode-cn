/*
@Time : 2019/9/30 10:41
@Author : zxr
@File : lengthOfLongestSubstring
@Software: GoLand
*/
package lettcode

import (
	"fmt"
	"math"
)

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
*/
func LengthOfLongestSubstring(str string) int {
	start, end := 0, 0
	maxLen := 0
	strMp := map[byte]int{}
	for end < len(str) {
		i, ok := strMp[str[end]]
		if ok {
			delete(strMp, str[end])
			start = int(math.Max(float64(start), float64(i)))
		}
		maxLen = int(math.Max(float64(maxLen), float64(end-start+1)))
		strMp[str[end]] = end + 1
		end++
	}
	return maxLen
}

func LengthOfLongestSubstring2(str string) int {
	if str == "" {
		return 0
	}
	start, end, res := 0, 0, 0
	lookup := map[byte]int{}
	for start < len(str) && end < len(str) {
		_, ok := lookup[str[end]]
		if !ok {
			lookup[str[end]] = end
			res = int(math.Max(float64(res), float64(end-start+1)))
			end += 1
		} else {
			delete(lookup, str[start])
			start += 1
		}
	}
	return res
}

func LengthOfLongestSubstring3(str string) int {
	lookup := map[byte]int{}
	start, res := 0, 0
	for i := range str {
		idx, ok := lookup[str[i]]
		if ok {
			start = int(math.Max(float64(start), float64(idx+1)))
		} else {
			start = int(math.Max(float64(start), float64(-1+1)))
		}
		fmt.Println(i, "--", start)
		res = int(math.Max(float64(res), float64(i-start+1)))
		lookup[str[i]] = i
	}
	return res
}
