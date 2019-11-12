/*
@Time : 2019/11/12 17:18
@Author : zxr
@File : 349
@Software: GoLand
*/
package _go

import "fmt"

/**
给定两个数组，编写一个函数来计算它们的交集。
示例 1:
输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2]
示例 2:
输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [9,4]
说明:

输出结果中的每个元素一定是唯一的。
我们可以不考虑输出结果的顺序。
*/
func Intersection(nums1 []int, nums2 []int) []int {
	var ret []int
	for _, v := range nums1 {
		j := 0
		var isEq bool
		for j < len(nums2) {
			if v == nums2[j] {
				isEq = true
				nums2[j] = -999
			}
			j++
		}
		if isEq {
			ret = append(ret, v)
		}
	}
	fmt.Println(ret)
	return ret
}
