/*
@Time : 2019/12/2 17:16
@Author : zxr
@File : 704
@Software: GoLand
*/
package _go

import (
	"fmt"
)

/**
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

示例 1:

输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
示例 2:

输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-search
**/
func Search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func Search2(nums []int, target int) int {
	fmt.Println("---------------------")
	fmt.Println(nums)
	var mid int
	if len(nums) == 0 {
		return -1
	}
	mid = len(nums) / 2
	if nums[mid] > target {
		return Search2(nums[0:mid], target)
	} else if nums[mid] < target {
		return Search2(nums[mid+1:], target)
	} else {
		return mid
	}
}
