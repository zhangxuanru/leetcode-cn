/*
@Time : 2019/12/3 15:12
@Author : zxr
@File : 383
@Software: GoLand
*/
package _go

import "strings"

/**
给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，判断第一个字符串ransom能不能由第二个字符串magazines里面的字符构成。如果可以构成，返回 true ；否则返回 false。
(题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。)
注意：

你可以假设两个字符串均只含有小写字母。

canConstruct("a", "b") -> false
canConstruct("aa", "ab") -> false
canConstruct("aa", "aab") -> true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ransom-note
**/

func CanConstruct(ransomNote string, magazine string) bool {
	mLen := len(magazine)
	rLen := len(ransomNote)
	mStrMap := make(map[byte]int)
	i := 0
	for i < mLen {
		mStrMap[magazine[i]]++
		i++
	}
	i = 0
	for i < rLen {
		if num, ok := mStrMap[ransomNote[i]]; ok && num >= 1 {
			mStrMap[ransomNote[i]]--
		} else {
			return false
		}
		i++
	}
	return true
}

//todo 推荐这种...
func canConstruct(ransomNote string, magazine string) bool {
	for _, v := range ransomNote {
		if strings.Count(ransomNote, string(v)) > strings.Count(magazine, string(v)) {
			return false
		}
	}
	return true
}
