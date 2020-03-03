/*
@Time : 2020/3/3 15:31
@Author : zxr
@File : 10.01.go
@Software: GoLand
*/
package go_interview

import (
	"fmt"
	"sort"
)

/**
给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。 编写一个方法，将 B 合并入 A 并排序。
初始化 A 和 B 的元素数量分别为 m 和 n。

示例:
输入:
A = [1,2,3,0,0,0], m = 3
B = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sorted-merge-lcci
**/
func Merge(A []int, m int, B []int, n int) {
	copy(A[m:m+n], B)
	sort.Ints(A)
}

func Merge2(A []int, m int, B []int, n int) {
	fmt.Println("A:", A)
	A = append(A[0:m], B...)
	fmt.Println(A)
}
