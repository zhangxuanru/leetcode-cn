package string

/**
给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

示例 1：
输入: s1 = "abc", s2 = "bca"
输出: true

示例 2：
输入: s1 = "abc", s2 = "bad"
输出: false

说明：
0 <= len(s1) <= 100
0 <= len(s2) <= 100
通过次数3,602提交次数5,108

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/check-permutation-lcci
*/
func CheckPermutation(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	s1AscNum := 0
	s2AscNum := 0
	for _, r := range []rune(s1) {
		s1AscNum += int(r)
	}
	for _, r := range []rune(s2) {
		s2AscNum += int(r)
	}
	return s1AscNum == s2AscNum
}

//推荐这种写法
func CheckPermutation2(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	l := len(s1)
	num1 := 0
	num2 := 0
	for i := 0; i < l; i++ {
		num1 += int(s1[i])
		num2 += int(s2[i])
	}
	return num1 == num2
}
