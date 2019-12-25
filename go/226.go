/*
@Time : 2019/12/25 16:40
@Author : zxr
@File : 226
@Software: GoLand
*/
package _go

/**
翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1
备注:

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/invert-binary-tree
 **/
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else {
		tmp := root.Left
		root.Left = root.Right
		root.Right = tmp
		InvertTree(root.Left)
		InvertTree(root.Right)
	}
	return root
}
