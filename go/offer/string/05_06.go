package string

/**
整数转换。编写一个函数，确定需要改变几个位才能将整数A转成整数B。

示例1:

 输入：A = 29 （或者0b11101）, B = 15（或者0b01111）
 输出：2
示例2:

 输入：A = 1，B = 2
 输出：2
提示:

A，B范围在[-2147483648, 2147483647]之间

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-integer-lcci
*/
func ConvertInteger(A int, B int) int {
	c := uint(A ^ B)
	cnt := 0
	for c != 0 {
		r := c & 1
		if r == 1 {
			cnt += 1
		}
		c >>= 1
	}
	return cnt
}

//fun2
func ConvertInteger2(A, B int) int {
	n := A ^ B
	count := 0
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	return count
}
