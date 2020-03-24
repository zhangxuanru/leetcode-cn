/*
@Time : 2020/3/24 14:24
@Author : zxr
@File : 57
@Software: GoLand
*/
package array

/***
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]

限制：

1 <= target <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof
*/

func FindContinuousSequence(target int) [][]int {
	j := 1
	sum := 0
	tmp := []int{}
	res := [][]int{}
	mid := target / 2
	for i := 1; i <= mid; i++ {
		sum += i
		tmp = append(tmp, i)
		if sum == target {
			res = append(res, tmp)
			tmp = []int{}
			sum = 0
			j++
			i = j - 1
			continue
		}
		if sum > target {
			j++
			i = j - 1
			sum = 0
			tmp = []int{}
			continue
		}
		if j+j+1 > target {
			break
		}
	}
	return res
}
