/*
@Time : 2019/11/23 15:53
@Author : zxr
@File : 169
@Software: GoLand
*/
package _go

import (
	"sort"
)

/**
给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在众数。

示例 1:

输入: [3,2,3]
输出: 3
示例 2:

输入: [2,2,1,1,1,2,2]
输出: 2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/majority-element
*/
func MajorityElement(nums []int) int {
	sort.Ints(nums)
	numLen := len(nums)
	modNum := numLen / 2
	count := 0
	majorNum := 0
	if numLen == 1 {
		return nums[0]
	}
	for i := 0; i < numLen-1; i++ {
		if nums[i] == nums[i+1] {
			count++
		} else {
			count = 0
		}
		if count >= modNum {
			majorNum = nums[i]
			break
		}
	}
	return majorNum
}

func majorityElement2(nums []int) int {
	var res int
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] > len(nums)/2 {
			res = v
			break
		}
	}
	return res
}
