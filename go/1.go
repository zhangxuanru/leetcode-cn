/*
@Time : 2019/9/10 10:38
@Author : zxr
@File : sum
@Software: GoLand
*/
package _go

import "fmt"

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
*/
func TwoSum(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				fmt.Println(nums[i], "+", nums[j], "=", target)
				return []int{i, j}
			}
		}
	}
	return []int{}
}

//以空间换时间，建议这种
func TwoSum2(nums []int, target int) []int {
	lookup := map[int]int{}
	for j, num := range nums {
		if i, ok := lookup[target-num]; ok {
			return []int{i, j}
		} else {
			lookup[num] = j
		}
	}

	//length := len(nums)
	//lookup := make(map[int]int)
	//for i := 0; i < length; i++ {
	//	if j, ok := lookup[target-nums[i]]; ok {
	//		return []int{j, i}
	//	} else {
	//		lookup[nums[i]] = i
	//	}
	//}
	return []int{}
}
