package array

/**
给定一个整数数组（有正数有负数），找出总和最大的连续数列，并返回总和。

示例：

输入： [-2,1,-3,4,-1,2,1,-5,4]
输出： 6
解释： 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶：

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contiguous-sequence-lcci
*/
func MaxSubArray(nums []int) int {
	max := -0x3f3f3f3f
	tmp := 0
	if len(nums) == 1 {
		return nums[0]
	}
	for _, val := range nums {
		tmp += val
		if tmp < val {
			tmp = val
		}
		if max < tmp {
			max = tmp
		}
	}
	return max
}
