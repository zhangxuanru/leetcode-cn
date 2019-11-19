/*
@Time : 2019/11/19 15:56
@Author : zxr
@File : 821
@Software: GoLand
*/
package _go

import "fmt"

/**
给定一个字符串 S 和一个字符 C。返回一个代表字符串 S 中每个字符到字符串 S 中的字符 C 的最短距离的数组。

示例 1:

输入: S = "loveleetcode", C = 'e'
输出: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]
说明:

字符串 S 的长度范围为 [1, 10000]。
C 是一个单字符，且保证是字符串 S 里的字符。
S 和 C 中的所有字母均为小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shortest-distance-to-a-character
*/
func ShortestToChar(S string, C byte) []int {
	sLen := len(S)
	nums := make([]int, sLen)
	for k, b := range []byte(S) {
		if b == C {
			nums[k] = 0
		} else {
			l := k - 1
			r := k + 1
			minNum := 0
			//从左边找
			for l >= 0 && S[l] != C {
				l--
			}
			//从右边找
			for r < sLen && S[r] != C {
				r++
			}
			//判断 两边的距离大小
			if r >= sLen {
				r = k
			}
			if l >= 0 && S[l] == C {
				minNum = k - l
			}
			fmt.Println("k=", k, "r=", r)
			if S[r] == C {
				if minNum == 0 || minNum > r-k {
					minNum = r - k
				}
			}
			nums[k] = minNum
		}
	}
	return nums
}
