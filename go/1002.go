/*
@Time : 2019/11/21 15:19
@Author : zxr
@File : 1002
@Software: GoLand
*/
package _go

import (
	"math"
	"strings"
)

/**
给定仅有小写字母组成的字符串数组 A，返回列表中的每个字符串中都显示的全部字符（包括重复字符）组成的列表。例如，如果一个字符在每个字符串中出现 3 次，但不是 4 次，则需要在最终答案中包含该字符 3 次。

你可以按任意顺序返回答案。

示例 1：

输入：["bella","label","roller"]
输出：["e","l","l"]
示例 2：

输入：["cool","lock","cook"]
输出：["c","o"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-common-characters
*/
func CommonChars(A []string) []string {
	var strRet []string
	var charLen = 26
	result := make([]int, charLen)
	for _, val := range A[0] {
		result[val-'a']++
	}
	for j := 1; j < len(A); j++ {
		tmp := make([]rune, charLen)
		for _, val := range A[j] {
			tmp[val-'a']++
		}
		for i := 0; i < charLen; i++ {
			result[i] = int(math.Min(float64(result[i]), float64(tmp[i])))
		}
	}
	for i := 0; i < charLen; i++ {
		if result[i] > 0 {
			count := result[i]
			for count > 0 {
				strRet = append(strRet, string(i+'a'))
				count--
			}
		}
	}
	return strRet
}

//方法二
func CommonChars2(A []string) []string {
	var strRet []string
	if len(A) < 1 {
		return A
	}
	strMp := make(map[rune]int)
	for _, val := range A[0] {
		strMp[val]++
	}
	for s, i := range strMp {
		str := string(s)
		for j := 1; j < len(A); j++ {
			count := strings.Count(A[j], str)
			if count < i {
				i = count
			}
		}
		if i == 0 {
			continue
		}
		for k := 0; k < i; k++ {
			strRet = append(strRet, str)
		}
	}
	return strRet
}
