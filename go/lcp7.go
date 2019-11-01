/*
@Time : 2019/9/12 18:20
@Author : zxr
@File : reverse
@Software: GoLand
*/
/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
示例 1:
输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
*/
package _go

func Reverse(x int) int {
	if x < 0 {
		return -Reverse(-x)
	}
	var res = 0
	for x != 0 {
		res = res*10 + x%10 //x%10 拿到最后的位数，之间是按10进一位
		x = x / 10
	}
	if res > 2147483647 || res < -2147483648 {
		return 0
	}
	return res
}
