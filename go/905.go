/*
@Time : 2019/11/11 20:58
@Author : zxr
@File : 905
@Software: GoLand
*/
package _go

/***
给定一个非负整数数组 A，返回一个数组，在该数组中， A 的所有偶数元素之后跟着所有奇数元素。

你可以返回满足此条件的任何数组作为答案。
示例：

输入：[3,1,2,4]
输出：[2,4,3,1]
输出 [4,2,3,1]，[2,4,1,3] 和 [4,2,1,3] 也会被接受。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-array-by-parity
*/

func SortArrayByParity(A []int) []int {
	l := 0
	r := len(A) - 1
	var evenNums []int
	var oddNums []int
	if r == 0 {
		return A
	}
	for l <= r {
		if A[l]%2 == 0 {
			evenNums = append(evenNums, A[l])
		} else {
			oddNums = append(oddNums, A[l])
		}
		l++
	}
	evenNums = append(evenNums, oddNums...)
	return evenNums
}

func sortArrayByParity2(A []int) []int {
	for left, right := 0, len(A)-1; left < right; {
		for left < right && A[left]%2 == 0 {
			left++
		}
		for right > left && A[right]%2 != 0 {
			right--
		}
		A[left], A[right] = A[right], A[left]
	}
	return A
}
