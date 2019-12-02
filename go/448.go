/*
@Time : 2019/11/28 12:13
@Author : zxr
@File : 448
@Software: GoLand
*/
package _go

import (
	"fmt"
	"math"
	"sort"
)

/**
给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
找到所有在 [1, n] 范围之间没有出现在数组中的数字。
您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。

示例:
输入:
[4,3,2,7,8,2,3,1]

输出:
[5,6]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array
**/

func FindDisappearedNumbers(nums []int) []int {
	var intRet []int
	for i := 0; i < len(nums); i++ {
		val := float64(nums[i])
		key := int(math.Abs(val) - 1) //找当前值的上一个值的位置
		if nums[key] > 0 {
			nums[key] = -nums[key]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			intRet = append(intRet, i+1)
		}
	}

	//	fmt.Println(nums)
	//fmt.Println(intRet)
	return intRet
}

func FindDisappearedNumbers_two(nums []int) []int {
	exists := make([]bool, len(nums)+len(nums))
	var intRet []int
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
		exists[v] = true
	}
	//todo 明天提交测试一下
	for i, v := range exists {
		if v == false && len(intRet) == 0 && i > max {
			intRet = append(intRet, i)
			break
		}
		if v == false && i > 0 && i < max {
			intRet = append(intRet, i)
			continue
		}
	}
	return intRet
}

func FindDisappearedNumbers4(nums []int) []int {
	var intRet []int
	if len(nums) == 0 {
		return intRet
	}
	sort.Ints(nums)
	j := 1
	i := 0
	tmp := 0
	fmt.Println(nums)
	for i = 0; i < len(nums) && j < len(nums)-1; i++ {
		fmt.Println("i=", i, "j=", j, "tmp=", tmp)
		if j > 0 && nums[j]-1 != nums[j-1] && nums[j] != nums[j-1] {
			intRet = append(intRet, nums[i]-1)
		}
		if nums[i]+1 != nums[j] && nums[i] != nums[j] {
			tmp = nums[i] + 1
			nums[i] = tmp
			intRet = append(intRet, tmp)

		} else {
			j++
		}

		//if nums[i]+1 == nums[j] {
		//	j++
		//	continue
		//}
		//if nums[i] == nums[j] && i != j {
		//	j++
		//	continue
		//}
		//tmp = nums[i] + 1
		//nums[i] = tmp
		//i--
		//intRet = append(intRet, tmp)
	}
	if len(intRet) == 0 {
		intRet = append(intRet, nums[i]+1)
	}

	fmt.Println(intRet)
	return intRet
}

func FindDisappearedNumbers3(nums []int) []int {
	var intRet []int
	if len(nums) == 0 {
		return intRet
	}
	sort.Ints(nums)
	j := 1
	tmp := 0
	for i := 0; i < len(nums) && j < len(nums); i++ {
		if nums[i]+1 == nums[j] || nums[i] == nums[j] {
			j++
			continue
		} else {
			tmp = nums[i] + 1
			nums[i] = tmp
			i--
			intRet = append(intRet, tmp)
		}
	}
	fmt.Println(intRet)
	return intRet
}

func FindDisappearedNumbers2(nums []int) []int {
	var intRet []int
	if len(nums) == 0 {
		return intRet
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	fmt.Println(nums)
	j := 1
	for i := 0; i < len(nums) && j < len(nums); i++ {
		if nums[i]-1 == nums[j] || nums[i] == nums[j] {
			j++
			continue
		} else {
			fmt.Println(nums[i], "--", nums[j])

			intRet = append(intRet, nums[i]-1)
			j++
		}
	}
	fmt.Println(intRet)
	return intRet
}
