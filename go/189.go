/*
@Time : 2019/12/24 12:16
@Author : zxr
@File : 189
@Software: GoLand
*/
package _go

import "fmt"

/**
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
示例 1:

输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]

解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-array
 **/

func Rotate(nums []int, k int) {
	length := len(nums)
	k %= length
	nums = append(nums[:0], append(nums[(length-k):], nums[0:(length-k)]...)...)

	return
	//l := len(nums) - k
	//if len(nums) == 1 {
	//	return
	//}
	////外层append的目的就是为了把新的切片追加到原数组指针上去
	//nums = append(nums[:0], append(nums[l:], nums[0:l]...)...)
}

//todo 超出时间限制
func Rotate2(nums []int, k int) {
	j := len(nums) - 1
	for k > 0 {
		tmp := nums[j]
		for i := j; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = tmp
		fmt.Println(nums)
		k--
	}
}
