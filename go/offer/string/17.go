/*
@Time : 2020/3/23 18:02
@Author : zxr
@File : 17
@Software: GoLand
*/
package string

/**
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

示例 1:

输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof
*/
func PrintNumbers(n int) []int {
	res := 0
	for i := 1; i <= n; i++ {
		res = res*10 + 9
	}
	ret := make([]int, res)
	for i := 0; i < res; i++ {
		ret[i] = i + 1
	}
	return ret
}
