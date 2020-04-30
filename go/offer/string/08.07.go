package string

import (
	"fmt"
)

/**
无重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合，字符串每个字符均不相同。

示例1:

 输入：S = "qwe"
 输出：["qwe", "qew", "wqe", "weq", "ewq", "eqw"]
示例2:

 输入：S = "ab"
 输出：["ab", "ba"]
*/
func Permutation(S string) []string {
	var strList = []string{S}

	return strList
}

//success
func Permutation2(S string) []string {
	n := len(S)
	ret := []string{}
	visited := map[int]bool{}
	var br func(now string, i int)

	br = func(now string, k int) {
		fmt.Println("params:", now, "i=", k)
		if len(now) == n {
			ret = append(ret, now)
			fmt.Println("now visited:", visited, "i=", k)
			return
		}
		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			br(now+string(S[i]), i)
			visited[i] = false
			fmt.Println("i=", i, "visited:", visited)
		}
	}
	br("", 0)
	return ret
}

//success
func Permutation3(str []rune, start int) {
	if str == nil {
		return
	}
	if start == len(str)-1 {
		fmt.Println(string(str))
	} else {
		for i := start; i < len(str); i++ {
			SwapRune(str, start, i)
			Permutation3(str, start+1)
			SwapRune(str, start, i)
		}
	}
}
func SwapRune(data []rune, x, y int) {
	tmp := data[x]
	data[x] = data[y]
	data[y] = tmp
}
