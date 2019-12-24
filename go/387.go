/*
@Time : 2019/12/23 18:45
@Author : zxr
@File : 387
@Software: GoLand
*/
package _go

/**
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string
**/
func FirstUniqChar(s string) int {
	sByteMap := make(map[rune]int)
	for _, v := range s {
		sByteMap[v]++
	}
	for k, v := range s {
		if n, ok := sByteMap[v]; ok && n == 1 {
			return k
		}
	}
	return -1
}

func FirstUniqChar2(s string) int {
	var intArr [26]int
	for _, v := range s {
		intArr[v-'a']++
	}
	for k, v := range s {
		if intArr[v-'a'] == 1 {
			return k
		}
	}
	return -1
}
