/*
@Time : 2019/11/18 16:37
@Author : zxr
@File : 867
@Software: GoLand
*/
package _go

/**
给定一个矩阵 A， 返回 A 的转置矩阵。

矩阵的转置是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。

示例 1：

输入：[[1,2,3],[4,5,6],[7,8,9]]
输出：[[1,4,7],[2,5,8],[3,6,9]]
示例 2：

输入：[[1,2,3],[4,5,6]]
输出：[[1,4],[2,5],[3,6]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/transpose-matrix
*/
func Transpose(A [][]int) [][]int {
	l := len(A)
	ml := len(A[0])
	var ret [][]int
	for i := 0; i < ml; i++ {
		var nums []int
		for k := 0; k < l; k++ {
			nums = append(nums, A[k][i])
		}
		ret = append(ret, nums)
	}
	return ret
}
