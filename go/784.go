/*
@Time : 2019/11/23 16:29
@Author : zxr
@File : 784
@Software: GoLand
*/
package _go

import "fmt"

/**
//todo 这道题不会

给定一个字符串S，通过将字符串S中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。

示例:
输入: S = "a1b2"
输出: ["a1b2", "a1B2", "A1b2", "A1B2"]

输入: S = "3z4"
输出: ["3z4", "3Z4"]

输入: S = "12345"
输出: ["12345"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-case-permutation
*/
func LetterCasePermutation(S string) []string {
	//fmt.Println("s=", S)
	//letter := make(map[rune]rune)
	//upper := make(map[rune]rune)
	//
	//var j int
	//var strRet = []string{S}
	//
	//for i := 'a'; i <= 'z'; i++ {
	//	letter[i] = i - 32
	//}
	//for i := 'A'; i <= 'Z'; i++ {
	//	upper[i] = i + 32
	//}
	//
	////
	////for {
	////	if j >= len(S) {
	////		break
	////	}
	////	rStr := rune(S[j])
	////	needStr := ""
	////	if v, ok := letter[rStr]; ok {
	////		needStr = string(S[0:j] + string(v) + S[j+1:])
	////	}
	////	if v, ok := upper[rStr]; ok {
	////		needStr = string(S[0:j] + string(v) + S[j+1:])
	////	}
	////	if len(needStr) > 0 {
	////		strRet = append(strRet, needStr)
	////		i := j + 1
	////		for i < len(needStr) {
	////			rStr := rune(needStr[i])
	////			if v, ok := letter[rStr]; ok {
	////				needStr = string(needStr[0:i] + string(v) + needStr[i+1:])
	////				strRet = append(strRet, needStr)
	////			}
	////			i++
	////		}
	////	}
	////	j++
	////}
	//
	//fmt.Println(strRet)
	return []string{"aaa"}
}

//todo 没完成， 周一继续
//https://blog.csdn.net/zpznba/article/details/88969618
func LetterCasePermutation2(S string) []string {
	letter := make(map[rune]rune)
	upper := make(map[rune]rune)
	var j int
	var strRet = []string{S}
	for i := 'a'; i <= 'z'; i++ {
		letter[i] = i - 32
	}
	for i := 'A'; i <= 'Z'; i++ {
		upper[i] = i + 32
	}
	for {
		if j >= len(S) {
			break
		}
		r := rune(S[j])
		var needR rune
		if v, ok := letter[r]; ok {
			needR = v
		}
		if v, ok := upper[r]; ok {
			needR = v
		}
		if needR > 0 {
			needStr := string(S[0:j]) + string(needR) + string(S[j+1:])
			strRet = append(strRet, needStr)
		}
		j++
	}

	fmt.Println(strRet)
	return strRet
}

//继续
func LetterCasePermutation3(S string) []string {
	charMap := make(map[int]byte)
	rStrMap := make(map[string]bool)
	rStrMap[S] = true
	var rStr []string
	sLen := len(S)
	for i := 0; i < sLen; i++ {
		if (S[i] >= 65 && S[i] <= 90) || (S[i] >= 97 && S[i] <= 122) {
			charMap[i] = S[i]
		}
	}
	if len(charMap) == 0 {
		rStr = append(rStr, S)
		return rStr
	}
	for k, v := range charMap {
		var lowerStr byte
		var upperStr byte
		var t1, t2, str1, str2 string
		if v >= 65 && v <= 90 {
			upperStr = v + 32
			lowerStr = v
		}
		if v >= 97 && v <= 122 {
			lowerStr = v - 32
			upperStr = v
		}
		if len(t1) == 0 {
			t1 = string(S[0:k]) + string(lowerStr) + S[k+1:]
			rStrMap[t1] = true
		}
		if len(t2) == 0 {
			t2 = string(S[0:k]) + string(upperStr) + S[k+1:]
			rStrMap[t2] = true
		}
		for key, val := range charMap {
			if key == k {
				continue
			}
			var lowerStr byte
			var upperStr byte
			if val >= 65 && val <= 90 {
				upperStr = val + 32
				lowerStr = val
			}
			if val >= 97 && val <= 122 {
				lowerStr = val - 32
				upperStr = val
			}
			t1 = string(t1[0:key]) + string(upperStr) + t1[key+1:]
			if len(str1) == 0 {
				str1 = string(t1[0:key]) + string(lowerStr) + t1[key+1:]
			} else {
				str1 = string(str1[0:key]) + string(lowerStr) + str1[key+1:]
			}

			t2 = string(t2[0:key]) + string(upperStr) + t2[key+1:]
			if len(str2) == 0 {
				str2 = string(t2[0:key]) + string(lowerStr) + t2[key+1:]
			} else {
				str2 = string(str2[0:key]) + string(lowerStr) + str2[key+1:]
			}
			s1 := string(S[0:key]) + string(upperStr) + S[key+1:]
			s2 := string(S[0:key]) + string(lowerStr) + S[key+1:]
			rStrMap[t1] = true
			rStrMap[str1] = true
			rStrMap[str2] = true
			rStrMap[t2] = true
			rStrMap[s1] = true
			rStrMap[s2] = true

		}
	}
	for str, _ := range rStrMap {
		rStr = append(rStr, str)
	}
	fmt.Println(rStr)
	return rStr
}
