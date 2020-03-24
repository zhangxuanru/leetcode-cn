package string

/**
给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。

回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。

回文串不一定是字典当中的单词。

示例1：

输入："tactcoa"
输出：true（排列有"tacocat"、"atcocta"，等等）

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-permutation-lcci
*/
func CanPermutePalindrome(s string) bool {
	sMap := make(map[rune]int)
	divNum := 0
	for _, r := range s {
		sMap[r]++
	}
	for _, num := range sMap {
		if num%2 != 0 {
			divNum++
		}
	}
	return divNum <= 1
}
