/*
@Time : 2019/11/26 18:41
@Author : zxr
@File : 1046
@Software: GoLand
*/
package _go

import (
	"fmt"
	"sort"
)

/**
有一堆石头，每块石头的重量都是正整数。
每一回合，从中选出两块最重的石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
如果 x == y，那么两块石头都会被完全粉碎；
如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/last-stone-weight
**/
func LastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Sort(sort.Reverse(sort.IntSlice(stones)))
		fmt.Println(stones)
		y := stones[0]
		x := stones[1]
		stones = stones[2:]
		if x != y {
			tmp := y - x
			stones = append(stones, tmp)
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}
