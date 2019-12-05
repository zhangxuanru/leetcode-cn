/*
@Time : 2019/12/5 10:30
@Author : zxr
@File : 888
@Software: GoLand
*/
package _go

import "fmt"

/**
爱丽丝和鲍勃有不同大小的糖果棒：A[i] 是爱丽丝拥有的第 i 块糖的大小，B[j] 是鲍勃拥有的第 j 块糖的大小。
因为他们是朋友，所以他们想交换一个糖果棒，这样交换后，他们都有相同的糖果总量。（一个人拥有的糖果总量是他们拥有的糖果棒大小的总和。）
返回一个整数数组 ans，其中 ans[0] 是爱丽丝必须交换的糖果棒的大小，ans[1] 是 Bob 必须交换的糖果棒的大小。
如果有多个答案，你可以返回其中任何一个。保证答案存在。

示例 1：

输入：A = [1,1], B = [2,2]
输出：[1,2]
示例 2：

输入：A = [1,2], B = [2,3]
输出：[1,2]
示例 3：

输入：A = [2], B = [1,3]
输出：[2,3]
示例 4：

输入：A = [1,2,5], B = [2,4]
输出：[5,4]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fair-candy-swap
 **/

func FairCandySwap(A []int, B []int) []int {
	var aSum, bSum int
	var ret []int
	bMapData := make(map[int]struct{})
	for _, v := range A {
		aSum += v
	}
	for _, v := range B {
		bSum += v
		bMapData[v] = struct{}{}
	}
	avg := (aSum + bSum) / 2

	fmt.Println("aSum:", aSum, " bSum:", bSum, " avg:", avg)

	for _, v := range A {
		b := avg + v - aSum
		fmt.Println("b:", b, " v:", v)
		if _, ok := bMapData[b]; ok {
			ret = []int{v, b}
			return ret
		}
	}
	return ret
}

func FairCandySwap2(A []int, B []int) []int {
	var aSum, bSum int
	aMap := make(map[int]struct{})
	for _, v := range A {
		aSum += v
		aMap[v] = struct{}{}
	}
	for _, v := range B {
		bSum += v
	}
	diff := (aSum+bSum)/2 - bSum
	for _, v := range B {
		key := v + diff
		if _, ok := aMap[key]; ok {
			return []int{v + diff, v}
		}
	}
	return []int{}
}
