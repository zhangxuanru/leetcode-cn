/*
@Time : 2019/11/18 17:01
@Author : zxr
@File : 961
@Software: GoLand
*/
package _go

/**
在大小为 2N 的数组 A 中有 N+1 个不同的元素，其中有一个元素重复了 N 次。

返回重复了 N 次的那个元素。

示例 1：

输入：[1,2,3,3]
输出：3
示例 2：

输入：[2,1,2,5,3,2]
输出：2
示例 3：

输入：[5,1,5,2,5,3,5,4]
输出：5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-repeated-element-in-size-2n-array
*/
func RepeatedNTimes(A []int) int {
	intLen := len(A)
	binNum := intLen / 2
	numMap := make(map[int]int)
	var num int
	for i := 0; i < intLen; i++ {
		key := A[i]
		numMap[key]++
		if numMap[key] == binNum {
			num = key
			break
		}
	}
	return num
}
