/*
@Time : 2019/10/30 16:19
@Author : zxr
@File : lcp1221
@Software: GoLand
*/
/**
在一个「平衡字符串」中，'L' 和 'R' 字符的数量是相同的。
给出一个平衡字符串 s，请你将它分割成尽可能多的平衡字符串。
返回可以通过分割得到的平衡字符串的最大数量。

示例 1：
输入：s = "RLRRLLRLRL"
输出：4
解释：s 可以分割为 "RL", "RRLL", "RL", "RL", 每个子字符串中都包含相同数量的 'L' 和 'R'。
示例 2：
输入：s = "RLLLLRRRLR"
输出：3
解释：s 可以分割为 "RL", "LLLRRR", "LR", 每个子字符串中都包含相同数量的 'L' 和 'R'。
示例 3：
输入：s = "LLLLRRRR"
输出：1
解释：s 只能保持原样 "LLLLRRRR".
链接：https://leetcode-cn.com/problems/split-a-string-in-balanced-strings
*/
package _go

import (
	"fmt"
	"strings"
)

func BalancedStringSplit(s string) int {
	l := len(s)
	index := 0
	var strList []string
	for index < l {
		val := s[index]
		for i := index + 1; i < l; i++ {
			if val != s[i] {
				diffLen := i - index
				countLen := i + diffLen
				if i+diffLen > l {
					countLen = l - 1
				}
				fmt.Println(string(val), "--index:", index, "--i:", i, "countLen:", countLen)
				//RLLRRRLLLR
				lastStr := s[index:countLen]
				if countLen >= l {
					countLen = l - 1
				}
				//if diffLen > 1 && s[countLen] == s[countLen-1] && countLen < l-1 {
				//	fmt.Println(countLen, string(s[countLen]), string(s[countLen-1]))
				//	index = i - 1
				//	break
				//}
				if strings.Count(lastStr, "L") == len(lastStr)/2 {
					strList = append(strList, lastStr)
					index = i + diffLen - 1
					break
				} else {
					index = i + diffLen - 1
					break
				}
			}
		}
		index++
	}
	fmt.Println(strList)
	return len(strList)
}

//ok
func BalancedStringSplit2(s string) int {
	lenL := 0
	lenR := 0
	k := 0
	for _, v := range s {
		if v == 'L' {
			lenL++
		}
		if v == 'R' {
			lenR++
		}
		if lenL == lenR {
			lenL = 0
			lenR = 0
			k = k + 1
		}
	}
	return k
}
