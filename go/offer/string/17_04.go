package string

import "fmt"

/**
数组nums包含从0到n的所有整数，但其中缺了一个。请编写代码找出那个缺失的整数。你有办法在O(n)时间内完成吗？

注意：本题相对书上原题稍作改动

示例 1：

输入：[3,0,1]
输出：2


示例 2：

输入：[9,6,4,2,3,5,7,0,1]
输出：8

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/missing-number-lcci
*/
//n*(n-1)/2
//n*(n+1)/2
func MissingNumber(nums []int) int {
	n := len(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return (n+1)*n/2 - sum
}

//test
func MissingNumberS(nums []int) {
	nums = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := len(nums)
	sum := n * (n - 1) / 2
	fmt.Println("sum:", sum)
}
