<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/9/10
 * Time: 11:31
 */
/*
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
 * */

function twoSum1($nums= array(),$target)
{
    $len = count($nums);
    for($i=0;$i<$len;$i++){
        for($j=$i+1;$j<$len;$j++){
            if($nums[$i] + $nums[$j] == $target){
                  return array($i,$j);
            }
        }
    }
    return array();
}

function twoSum2($nums = array(),$target)
{
    $lookUp = array();
   foreach($nums as $key => $val){
       if(isset($lookUp[$target - $val])){
           return array($lookUp[$target - $val],$key);
       }else{
           $lookUp[$val] = $key;
       }
   }
    return array();
}

$nums = array(2, 7, 11, 15, 20, 45, 67, 35, 78);
$target = 35;
$r = twoSum1($nums,$target);
print_r($r);
$r = twoSum2($nums,$target);
print_r($r);