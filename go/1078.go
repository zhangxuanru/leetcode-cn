/*
@Time : 2019/11/25 21:07
@Author : zxr
@File : 1078
@Software: GoLand
*/
package _go

import (
	"fmt"
	"strings"
)

/**
对于每种这样的情况，将第三个词 "third" 添加到答案中，并返回答案。

示例 1：

输入：text = "alice is a good girl she is a good student", first = "a", second = "good"
输出：["girl","student"]
示例 2：

输入：text = "we will we will rock you", first = "we", second = "will"
输出：["we","rock"]

提示：
1 <= text.length <= 1000
text 由一些用空格分隔的单词组成，每个单词都由小写英文字母组成
1 <= first.length, second.length <= 10
first 和 second 由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/occurrences-after-bigram
*/

func FindOcurrences(text string, first string, second string) []string {
	split := strings.Split(text, " ")
	var retStr []string
	i := 0
	for i < len(split)-2 {
		val := split[i]
		if val == first && split[i+1] == second {
			retStr = append(retStr, split[i+2])
			i += 2
			continue
		}
		i++
	}
	fmt.Println(retStr)
	return retStr
}
