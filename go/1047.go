/*
@Time : 2019/11/22 17:31
@Author : zxr
@File : 1047
@Software: GoLand
*/
package _go

import (
	"fmt"
	"strings"
)

/**
给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。

在 S 上反复执行重复项删除操作，直到无法继续删除。

在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

示例：

输入："abbaca"
输出："ca"
解释：
例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。之后我们得到字符串 "aaca"，其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string
*/

func RemoveDuplicates2(S string) string {
	for {
		tmpS := S
		for i := 0; i < len(S)-1; i++ {
			if S[i] == S[i+1] {
				tmpStr := string(S[i]) + string(S[i+1])
				S = strings.Replace(S, string(tmpStr), "", 1)
				continue
			}
		}
		if tmpS == S {
			break
		}
	}
	return S
}

//推荐这种
func RemoveDuplicates4(S string) string {
	var stack []byte
	stack = make([]byte, len(S))
	index := 1
	stack[0] = S[0]
	for i := 1; i < len(S); i++ {
		stack[index] = S[i]
		if index > 0 && stack[index] == stack[index-1] {
			index--
			continue
		}
		index++
	}

	fmt.Println(index)
	fmt.Println(stack[:index])
	return string(stack[:index])
}
