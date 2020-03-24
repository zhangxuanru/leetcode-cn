/*
@Time : 2020/3/24 16:00
@Author : zxr
@File : 39
@Software: GoLand
*/
package array

/**
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1:

输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
输出: 2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof
*/
func MajorityElement(nums []int) int {
	n := len(nums)
	midLen := n / 2
	numMap := make(map[int]int)
	for i := 0; i < n; i++ {
		if n, _ := numMap[nums[i]]; n >= midLen {
			return nums[i]
		}
		numMap[nums[i]]++
	}
	return -1
}

// func 2
func MajorityElement2(nums []int) int {
	count := 1
	val := nums[0]
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			val = nums[i]
		}
		if val == nums[i] {
			count++
		} else {
			count--
		}
	}
	return val
}
