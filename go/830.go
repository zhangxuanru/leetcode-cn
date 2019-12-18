/*
@Time : 2019/12/17 16:09
@Author : zxr
@File : 830
@Software: GoLand
*/
package _go

/**
在一个由小写字母构成的字符串 S 中，包含由一些连续的相同字符所构成的分组。
例如，在字符串 S = "abbxxxxzyy" 中，就含有 "a", "bb", "xxxx", "z" 和 "yy" 这样的一些分组。
我们称所有包含大于或等于三个连续字符的分组为较大分组。找到每一个较大分组的起始和终止位置。
最终结果按照字典顺序输出。
示例 1:
输入: "abbxxxxzzy"
输出: [[3,6]]
解释: "xxxx" 是一个起始于 3 且终止于 6 的较大分组。
示例 2:

输入: "abc"
输出: []
解释: "a","b" 和 "c" 均不是符合要求的较大分组。
示例 3:

输入: "abcdddeeeeaabbbcd"
输出: [[3,5],[6,9],[12,14]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/positions-of-large-groups
**/

func LargeGroupPositions(S string) [][]int {
	var ret [][]int
	i := 0
	sLen := len(S)
	var startS byte

	startIndex := 0
	endIndex := 0
	for i < sLen {
		c := S[i]
		if startS == ' ' {
			startS = c
			startIndex = i
			endIndex = i
			continue
		}
		if c == startS {
			endIndex++
		} else {
			if endIndex-startIndex >= 2 {
				ret = append(ret, []int{startIndex, endIndex})
			}
			startS = c
			startIndex = i
			endIndex = i
		}
		i++
	}
	if endIndex-startIndex >= 2 {
		ret = append(ret, []int{startIndex, endIndex})
	}
	return ret
}

//fun2
func LargeGroupPositions2(S string) [][]int {
	var ret [][]int
	i := 0
	sLen := len(S)
	var startS byte
	startIndex := 0
	endIndex := 0
	for i < sLen {
		c := S[i]
		if startS == ' ' {
			startS = c
			startIndex = i
			endIndex = i
			continue
		}
		if c == startS {
			endIndex++
		} else {
			startS = c
			endIndex = i
			startIndex = i
		}
		//fmt.Println("i:", i, "sLen:", sLen, string(c), ":", startIndex, "--", endIndex)
		if endIndex-startIndex >= 2 {
			if i == sLen-1 || S[i+1] != c {
				ret = append(ret, []int{startIndex, endIndex})
			}
		}
		i++
	}
	return ret
}
