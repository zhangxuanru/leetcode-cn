/*
@Time : 2019/11/28 11:51
@Author : zxr
@File : 485
@Software: GoLand
*/
package _go

/**
给定一个二进制数组， 计算其中最大连续1的个数。

示例 1:

输入: [1,1,0,1,1,1]
输出: 3
解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
注意：

输入的数组只包含 0 和1。
输入数组的长度是正整数，且不超过 10,000。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/max-consecutive-ones
 **/
func FindMaxConsecutiveOnes(nums []int) int {
	i := 0
	maxNum := 0
	tmpMax := 0
	for i < len(nums) {
		if nums[i] == 1 {
			tmpMax++
		}
		if tmpMax > maxNum {
			maxNum = tmpMax
		}
		if nums[i] == 0 {
			tmpMax = 0
		}
		i++
	}
	return maxNum
}
