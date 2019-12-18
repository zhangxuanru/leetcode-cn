<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/12/17
 * Time: 11:01
 */
/**
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
示例 1:
s = "abc", t = "ahbgdc"
返回 true.
示例 2:
s = "axc", t = "ahbgdc"
返回 false.
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/is-subsequence
 **/
class Solution {

    /**
     * @param String $s
     * @param String $t
     * @return Boolean
     */
    function isSubsequence($s, $t) {
            $i = 0;
            $j = 0;
            $sLen = strlen($s);
            $tLen = strlen($t);
            while($i < $sLen && $j < $tLen){
                if($s[$i] == $t[$j]){
                  $i++;
                }
             $j++;
         }
        return $i == $sLen;
    }
}

$class = new Solution();
$s = "abc";
$t = "ahbgdc";
$ret = $class->isSubsequence($s,$t);
var_dump($ret);
