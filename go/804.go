/*
@Time : 2019/11/1 11:28
@Author : zxr
@File : 804
@Software: GoLand
*/
/**
国际摩尔斯密码定义一种标准编码方式，将每个字母对应于一个由一系列点和短线组成的字符串， 比如: "a" 对应 ".-", "b" 对应 "-...", "c" 对应 "-.-.", 等等。
为了方便，所有26个英文字母对应摩尔斯密码表如下：
[".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
给定一个单词列表，每个单词可以写成每个字母对应摩尔斯密码的组合。例如，"cab" 可以写成 "-.-..--..."，(即 "-.-." + "-..." + ".-"字符串的结合)。我们将这样一个连接过程称作单词翻译。
返回我们可以获得所有词不同单词翻译的数量。
例如:
输入: words = ["gin", "zen", "gig", "msg"]
输出: 2
解释:
各单词翻译如下:
"gin" -> "--...-."
"zen" -> "--...-."
"gig" -> "--...--."
"msg" -> "--...--."
共有 2 种不同翻译, "--...-." 和 "--...--.".
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-morse-code-words
*/
package _go

func UniqueMorseRepresentations(words []string) int {
	sysMap := make(map[rune]string)
	wordSymMap := make(map[string]struct{})
	listSymbols := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	i := 0
	for k := 'a'; k <= 'z'; k++ {
		sysMap[k] = listSymbols[i]
		i++
	}
	for _, word := range words {
		symbolStr := ""
		for _, r := range word {
			if symbol, ok := sysMap[r]; ok {
				symbolStr += symbol
			}
		}
		if _, ok := wordSymMap[symbolStr]; !ok {
			wordSymMap[symbolStr] = struct{}{}
		}
	}
	return len(wordSymMap)
}

func UniqueMorseRepresentations2(words []string) int {
	wordSymMap := make(map[string]struct{})
	listSymbols := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	for _, word := range words {
		symbolStr := ""
		for _, r := range word {
			symbolStr += listSymbols[int(r-'a')]
		}
		if _, ok := wordSymMap[symbolStr]; !ok {
			wordSymMap[symbolStr] = struct{}{}
		}
	}
	return len(wordSymMap)
}
