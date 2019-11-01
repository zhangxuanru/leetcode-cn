/*
@Time : 2019/10/16 11:49
@Author : zxr
@File : threeSum
@Software: GoLand
*/
package _go

import "sort"

/*
 * 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 *
 * */
func ThreeSum(nums []int) [][]int {
	count := len(nums)
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			for k := j + 1; k < count; k++ {
				flag := true
				res := make([]int, 0)
				sum := nums[i] + nums[j] + nums[k]
				if sum == 0 {
					res = []int{nums[i], nums[j], nums[k]}
					for _, val := range result {
						if val[0] == res[0] && val[1] == res[1] && val[2] == res[2] {
							flag = false
							break
						}
					}
				}
				if flag == true && sum == 0 {
					result = append(result, res)
				}
			}
		}
	}
	return result
}

/**
推荐这种方法
*/
func ThreeSum2(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	result := make([][]int, 0)
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := length - 1
		for l < r {
			temp := nums[i] + nums[l] + nums[r]
			if temp == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
			if temp > 0 {
				r--
			}
			if temp < 0 {
				l++
			}
		}
	}

	return result
}
