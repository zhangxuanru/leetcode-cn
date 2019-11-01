/*
@Time : 2019/11/1 18:27
@Author : zxr
@File : 104
@Software: GoLand
*/
/**
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。
示例：
给定二叉树 [3,9,20,null,null,15,7]，
    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree
*/
package _go

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Gen(root, 1)
}

func Gen(root *TreeNode, num int) int {
	var l, r int
	if root == nil {
		return num
	}
	if root.Left != nil {
		l = Gen(root.Left, num+1)
	}
	if root.Right != nil {
		r = Gen(root.Right, num+1)
	}
	if l > num {
		num = l
	}
	if r > num {
		num = r
	}
	return num
}

func MaxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.Max(float64(MaxDepth2(root.Left)), float64(MaxDepth2(root.Right))) + 1
	return int(max)
}
