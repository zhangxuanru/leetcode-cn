/*
@Time : 2020/3/24 15:24
@Author : zxr
@File : 03
@Software: GoLand
*/
package array

/**
找出数组中重复的数字。
在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
*/
func FindRepeatNumber(nums []int) int {
	n := len(nums)
	r := n - 1
	var ret int
	numMap := make(map[int]struct{})
	for i := 0; i < n; i++ {
		if _, ok := numMap[nums[i]]; ok {
			ret = nums[i]
			break
		}
		if _, ok1 := numMap[nums[r]]; ok1 {
			ret = nums[r]
			break
		}
		if nums[i] == nums[r] || nums[i] == nums[i+1] {
			ret = nums[i]
			break
		}
		numMap[nums[i]] = struct{}{}
		numMap[nums[r]] = struct{}{}
		r--
	}
	return ret
}
