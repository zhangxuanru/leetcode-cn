/*
@Time : 2019/11/28 19:01
@Author : zxr
@File : 268
@Software: GoLand
*/
package _go

/**
给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。

示例 1:

输入: [3,0,1]
输出: 2
示例 2:

输入: [9,6,4,2,3,5,7,0,1]
输出: 8
说明:
你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/missing-number
 **/
func MissingNumber(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	realSum := (len(nums) * (len(nums) + 1)) / 2
	return realSum - sum
}

func MissingNumber2(nums []int) int {
	exists := make([]bool, len(nums)+1)
	for _, v := range nums {
		exists[v] = true
	}
	for i, v := range exists {
		if v == false {
			return i
		}
	}
	return -1
}
