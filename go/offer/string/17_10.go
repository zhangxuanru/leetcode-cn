package string

/**
如果数组中多一半的数都是同一个，则称之为主要元素。给定一个整数数组，
找到它的主要元素。若没有，返回-1。

示例 1：

输入：[1,2,5,9,5,9,5,5,5]
输出：5


示例 2：

输入：[3,2]
输出：-1


示例 3：

输入：[2,2,1,1,1,2,2]
输出：2


说明：
你有办法在时间复杂度为 O(N)，空间复杂度为 O(1) 内完成吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-majority-element-lcci
*/
func MajorityElement(nums []int) int {
	midNum := len(nums)/2 + 1
	numMap := make(map[int]int)
	for _, val := range nums {
		numMap[val]++
		if numMap[val] >= midNum {
			return val
		}
	}
	return -1
}

//fun2
func MajorityElement2(nums []int) int {
	cnt := 0
	ret := 0
	for _, val := range nums {
		if cnt == 0 {
			ret = val
			cnt++
		}
		if val == ret {
			cnt++
		} else {
			cnt--
		}
	}
	if cnt > 0 {
		return ret
	}
	return -1
}
