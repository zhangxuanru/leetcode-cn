/*
@Time : 2019/12/26 16:12
@Author : zxr
@File : 283
@Software: GoLand
*/
package _go

import "fmt"

/**
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/move-zeroes
 **/
func MoveZeroes(nums []int) {
	length := len(nums)
	i := 0
	j := 0
	for i < length {
		if nums[i] != 0 {
			i++
			continue
		}
		for j = i; j < length; j++ {
			if nums[j] != 0 {
				break
			}
		}
		fmt.Println("i=", i, "j=", j)
		if j >= length || nums[j] == 0 {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
	}
}

func MoveZeroes2(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 && i != j {
			nums[i] = nums[j]
			i++
		}
	}
	for i < len(nums) {
		nums[i] = 0
		i++
	}
}
