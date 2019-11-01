/*
@Time : 2019/10/31 18:27
@Author : zxr
@File : lcp461
@Software: GoLand
*/
/**
两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
给出两个整数 x 和 y，计算它们之间的汉明距离。
注意：
0 ≤ x, y < 231.
示例:
输入: x = 1, y = 4
输出: 2
解释:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/hamming-distance
*/
package _go

func HammingDistance(x int, y int) int {
	var num int
	for x > 0 || y > 0 {
		if x%2 != y%2 {
			num++
		}
		x = x / 2
		y = y / 2
	}
	return num
}
