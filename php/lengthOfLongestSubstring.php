<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/10/16
 * Time: 10:44
 */

/**
 * 给定一个字符串，找出其中不含重复字符的最长子串的长度
 * @param $str
 * @return int
 */
function lengthOfLongestSubstring($str)
{
    $arr = array();
    $start = 0;
    $res = 0;
    $i = 0;
    $len = strlen($str);
   while($i<$len){
       $s = $str[$i];
        if(isset($arr[$s])){
            $idx = $arr[$s];
            $start = max($start,$idx+1);
        }
        $res = max($res,$i-$start+1);
        $arr[$s] = $i;
        $i++;
    }
    return $res;
}

$str="abcabcbb";
$r = lengthOfLongestSubstring($str);
var_dump($r);