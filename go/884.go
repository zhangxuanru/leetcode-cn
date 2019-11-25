/*
@Time : 2019/11/25 19:10
@Author : zxr
@File : 884
@Software: GoLand
*/
package _go

import (
	"fmt"
	"strings"
)

/**
给定两个句子 A 和 B 。 （句子是一串由空格分隔的单词。每个单词仅由小写字母组成。）
如果一个单词在其中一个句子中只出现一次，在另一个句子中却没有出现，那么这个单词就是不常见的。
返回所有不常用单词的列表。
您可以按任何顺序返回列表。

示例 1：
输入：A = "this apple is sweet", B = "this apple is sour"
输出：["sweet","sour"]
示例 2：

输入：A = "apple apple", B = "banana"
输出：["banana"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/uncommon-words-from-two-sentences
*/

func UncommonFromSentences(A string, B string) []string {
	splitA := strings.Split(A, " ")
	splitB := strings.Split(B, " ")
	aMap := make(map[string]int)
	bMap := make(map[string]int)
	var retStr []string
	for _, chars := range splitA {
		aMap[chars]++
	}
	for _, chars := range splitB {
		bMap[chars]++
	}
	for chars, num := range aMap {
		if num > 1 || bMap[chars] > 0 {
			continue
		}
		retStr = append(retStr, chars)
	}
	for chars, num := range bMap {
		if num > 1 || aMap[chars] > 0 {
			continue
		}
		retStr = append(retStr, chars)
	}
	fmt.Println(retStr)
	return retStr
}

func UncommonFromSentences2(A string, B string) []string {
	splitA := strings.Split(A, " ")
	splitB := strings.Split(B, " ")
	var strRet []string
	splitA = append(splitA, splitB...)
	strMap := make(map[string]int)
	for _, str := range splitA {
		strMap[str]++
	}
	for str, num := range strMap {
		if num > 1 {
			continue
		}
		strRet = append(strRet, str)
	}
	fmt.Println(strRet)
	return strRet
}
