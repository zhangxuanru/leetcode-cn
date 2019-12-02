/*
@Time : 2019/12/2 21:13
@Author : zxr
@File : 167
@Software: GoLand
*/
package _go

/**
给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。

函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。

说明:

返回的下标值（index1 和 index2）不是从零开始的。
你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
示例:

输入: numbers = [2, 7, 11, 15], target = 9
输出: [1,2]
解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted
**/

func TwoSum8(numbers []int, target int) []int {
	dataMp := make(map[int]int)
	var ret []int
	for k, v := range numbers {
		key := target - v
		if index, ok := dataMp[key]; ok {
			if index > k+1 {
				ret = append(ret, k+1, index)
			} else {
				ret = append(ret, index, k+1)
			}
		}
		dataMp[v] = k + 1
	}
	return ret
}

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for {
		low := numbers[i]
		high := numbers[j]
		if low+high > target {
			j = j - 1
		} else if low+high < target {
			i = i + 1
		} else {
			break
		}
		if j <= i {
			return []int{}
		}
	}
	return []int{i + 1, j + 1}
}
