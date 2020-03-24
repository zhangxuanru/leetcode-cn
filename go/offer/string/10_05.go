package string

/**
稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。

示例1:

 输入: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ta"
 输出：-1
 说明: 不存在返回-1。
示例2:

 输入：words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ball"
 输出：4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sparse-array-search-lcci
*/

func FindString(words []string, s string) int {
	l := 0
	r := len(words) - 1
	for l <= r {
		mid := (l + r) / 2
		tm := mid + 1
		for words[mid] == "" {
			mid--
		}
		if mid < l {
			mid = tm
		}
		if words[mid] == s {
			return mid
		}
		if words[mid] < s {
			l = mid + 1
		}
		if words[mid] > s {
			r = mid - 1
		}
	}
	return -1
}
