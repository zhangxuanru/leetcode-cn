<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/12/24
 * Time: 20:35
 */
/**
给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。
示例 1:
输入: [3, 2, 1]
 *
输出: 1
解释: 第三大的数是 1.
 *
示例 2:
输入: [1, 2]
输出: 2
解释: 第三大的数不存在, 所以返回最大的数 2 .
 *
示例 3:
输入: [2, 2, 3, 1]
输出: 1

解释: 注意，要求返回第三大的数，是指第三大且唯一出现的数。
存在两个值为2的数，它们都排第二。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/third-maximum-number
 **/
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function thirdMax($nums) {
            $max1 = PHP_INT_MIN;
            $max2 = PHP_INT_MIN;
            $max3 = PHP_INT_MIN;
        foreach($nums  as $item){
            if($item > $max1){
                $max3 = $max2;
                $max2 = $max1;
                $max1 = $item;
                continue;
            }
            if($item > $max2 && $item < $max1){
                 $max3 = $max2;
                 $max2 = $item;
                continue;
            }
            if($item < $max2 && $item > $max3){
                $max3 = $item;
            }
        }
        if($max3 > PHP_INT_MIN){
            return $max3;
        }
        return $max1;
    }
}

$class = new Solution();
$nums =[3, 2, 1];
$val = $class->thirdMax($nums);
echo $val;
