/*
@Time : 2019/11/25 18:08
@Author : zxr
@File : 389
@Software: GoLand
*/
package _go

import "fmt"

/**
给定两个字符串 s 和 t，它们只包含小写字母。
字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
请找出在 t 中被添加的字母。

示例:

输入：
s = "abcd"
t = "abcde"

输出：
e

解释：
'e' 是那个被添加的字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-the-difference
*/
func FindTheDifference(s string, t string) byte {
	rMap := make(map[rune]int)
	tMap := make(map[rune]int)
	var r byte
	for _, v := range s {
		rMap[v]++
	}
	for _, v := range t {
		tMap[v]++
	}
	for k, v := range tMap {
		if v != rMap[k] {
			r = byte(k)
			break
		}
	}
	return r
}

//todo 别人的写法，这种写法更优秀
func FindTheDifference2(s string, t string) byte {
	// 本质上还是求解single number的问题
	var c byte = 0
	for _, ch := range s {
		c ^= byte(ch)
	}
	for _, ch := range t {
		c ^= byte(ch)
	}
	fmt.Println("c=", c)
	return c
}
