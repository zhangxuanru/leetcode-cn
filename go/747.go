/*
@Time : 2019/12/24 17:30
@Author : zxr
@File : 747
@Software: GoLand
*/
package _go

import "fmt"

/**
在一个给定的数组nums中，总是存在一个最大元素 。

查找数组中的最大元素是否至少是数组中每个其他数字的两倍。

如果是，则返回最大元素的索引，否则返回-1。
示例 1:

输入: nums = [3, 6, 1, 0]
输出: 1
解释: 6是最大的整数, 对于数组中的其他整数,
6大于数组中其他元素的两倍。6的索引是1, 所以我们返回1.


示例 2:

输入: nums = [1, 2, 3, 4]
输出: -1
解释: 4没有超过3的两倍大, 所以我们返回 -1.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/largest-number-at-least-twice-of-others
 **/
func DominantIndex(nums []int) int {
	max := 0
	i := 0
	for k, v := range nums {
		if v > max {
			i = k
			max = v
		}
	}
	for _, v := range nums {
		fmt.Println("max:", max, "v=", v*2)
		if v != max && max < v*2 {
			return -1
			break
		}
	}
	return i
}
