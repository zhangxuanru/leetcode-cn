<?php

/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

*/
function isPalindrome($x){
    if($x < 0 ||($x>0 && $x%10 ==0) ){
        return false;
    }
    $res = 0;
    $y = $x;
    while($x!=0){
        $res = $res*10+$x%10;
        $x = intval($x/10);
    }
    return $res == $y;
}
$x=121;
$rs = isPalindrome($x);
