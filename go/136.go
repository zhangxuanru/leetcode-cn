/*
@Time : 2019/11/22 16:07
@Author : zxr
@File : 136
@Software: GoLand
*/
package _go

import (
	"sort"
)

/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例 2:

输入: [4,1,2,1,2]
输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-number
*/

func SingleNumber(nums []int) int {
	sort.Ints(nums)
	numLen := len(nums)
	var retNumber int
	if numLen == 1 {
		return nums[0]
	}
	for j := 1; j < numLen; j++ {
		r := j + 1
		l := j - 1
		if l == 0 && nums[j] != nums[l] {
			retNumber = nums[l]
			break
		}
		if nums[j] == nums[l] {
			continue
		}
		if r >= numLen {
			r = j
		}
		if nums[j] == nums[r] && r != j {
			continue
		}
		retNumber = nums[j]
		break
	}
	return retNumber
}
