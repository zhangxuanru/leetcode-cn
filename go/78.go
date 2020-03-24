/*
@Time : 2020/1/14 16:41
@Author : zxr
@File : 78
@Software: GoLand
*/
package _go

/**
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
说明：解集不能包含重复的子集。
示例:
输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
**/

//error  明天重写
func Subsets(nums []int) [][]int {
	var numsResult [][]int
	count := len(nums)
	for i := 0; i < count; i++ {
		var tmp []int
		for j := i; j < count; j++ {
			tmp = append(tmp, nums[j])
			numsResult = append(numsResult, tmp)
		}
	}
	return numsResult
}

//ok
func Subsets2(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}
	res := Subsets2(nums[1:])
	ll := len(res)
	for i := 0; i < ll; i++ {
		res = append(res, append([]int{nums[0]}, res[i]...))
	}
	return res
}
