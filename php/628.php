<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/12/8
 * Time: 17:46
 */
/**
给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。

示例 1:

输入: [1,2,3]
输出: 6
示例 2:

输入: [1,2,3,4]
输出: 24
注意:

给定的整型数组长度范围是[3,104]，数组中所有的元素范围是[-1000, 1000]。
输入的数组中任意三个数的乘积不会超出32位有符号整数的范围。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-product-of-three-numbers
  **/
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function maximumProduct($nums) {
        $max = 0;
        $len = count($nums);
        if($len == 0){
            return $max;
        }
        rsort($nums);
        $max = $nums[0] * $nums[1] * $nums[2];
        $max1 = $nums[0] * $nums[$len-1] * $nums[$len-2];
        $max = max($max,$max1);
        return $max;
    }
}

$class = new Solution();
$arr = [1,2,3,4];
$arr = [-4,-3,-2,-1,60];
$num = $class->maximumProduct($arr);
echo $num;
