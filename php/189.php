<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/12/24
 * Time: 15:05
 */
/**
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-array
 **/
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return NULL
     */
    function rotate(&$nums, $k) {
        while($k > 0){
            $val = array_pop($nums);
            array_unshift($nums,$val);
            $k--;
        }
    }

    function rotate2(&$nums, $k){
        $len = count($nums);
        $nums = array_merge(array_slice($nums, $len-$k, $k), array_slice($nums, 0, $len-$k));
    }
}

$class = new Solution();
$nums = [1,2,3,4,5,6,7];
$k = 3;

$class->rotate2($nums,$k);
print_r($nums);