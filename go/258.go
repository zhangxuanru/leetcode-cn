/*
@Time : 2019/11/21 18:33
@Author : zxr
@File : 258
@Software: GoLand
*/
package _go

import "fmt"

/**
给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。

示例:

输入: 38
输出: 2
解释: 各位相加的过程为：3 + 8 = 11, 1 + 1 = 2。 由于 2 是一位数，所以返回 2。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-digits
*/
func AddDigits(num int) int {
	//r := 0
	//for num > 0 {
	//	lastNum := num % 10
	//	num = num / 10
	//	r = r*10 + lastNum
	//}
	for num >= 10 {
		lastNum := num % 10
		num = num / 10
		num = lastNum + num
	}
	fmt.Println(num)
	return num
}
