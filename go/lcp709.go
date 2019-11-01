/*
@Time : 2019/10/31 16:22
@Author : zxr
@File : lcp709
@Software: GoLand
*/
/**
实现函数 ToLowerCase()，该函数接收一个字符串参数 str，并将该字符串中的大写字母转换成小写字母，之后返回新的字符串。

示例 1：
输入: "Hello"
输出: "hello"
示例 2：
输入: "here"
输出: "here"
示例 3：
输入: "LOVELY"
输出: "lovely"
*/
package _go

import (
	"bytes"
	"fmt"
	"unicode"
)

func ToLowerCase(str string) string {
	upMap := make(map[string]string)
	var byt bytes.Buffer
	for i := 'A'; i <= 'Z'; i++ {
		upperKey := fmt.Sprintf("%c", i)
		lowerVal := fmt.Sprintf("%c", i+32)
		upMap[upperKey] = lowerVal
	}
	for _, v := range str {
		if lower, ok := upMap[string(v)]; ok {
			byt.WriteString(lower)
		} else {
			byt.WriteRune(v)
		}
	}
	return byt.String()
}

func ToLowerCase2(str string) string {
	r := []rune(str)
	str = ""
	for i, v := range r {
		if unicode.IsUpper(v) {
			r[i] = unicode.ToLower(v)
		}
	}
	return string(r)
}

func ToLowerCase3(str string) string {
	upMap := make(map[rune]rune)
	for i := 'A'; i <= 'Z'; i++ {
		upMap[i] = i + 32
	}
	runeStr := []rune(str)
	str = ""
	for i, v := range runeStr {
		if lower, ok := upMap[v]; ok {
			runeStr[i] = rune(lower)
		}
	}
	return string(runeStr)
}
