/*
@Time : 2019/11/1 17:33
@Author : zxr
@File : 1051
@Software: GoLand
*/
/**
学校在拍年度纪念照时，一般要求学生按照 非递减 的高度顺序排列。

请你返回至少有多少个学生没有站在正确位置数量。该人数指的是：能让所有学生以 非递减 高度排列的必要移动人数。

示例：

输入：[1,1,4,2,1,3]
输出：3
解释：
高度为 4、3 和最后一个 1 的学生，没有站在正确的位置。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/height-checker
*/
package _go

import (
	"fmt"
	"sort"
)

func HeightChecker(heights []int) int {
	i := 0
	l := len(heights)
	for j := 0; j < l-1; j++ {
		if heights[j] < 0 {
			continue
		}
		for k := j + 1; k < l; k++ {
			fmt.Println("j=", heights[j], "k=", heights[k])
			if heights[j] > heights[k] && heights[k] > 0 {
				i++
				heights[k] = -1
			}
		}
	}
	return i
}

//ok
func HeightChecker2(heights []int) int {
	tmp := make([]int, len(heights))
	r := 0
	copy(tmp, heights)
	sort.Ints(heights)
	for i, v := range tmp {
		if v != heights[i] {
			r++
		}
	}
	return r
}
