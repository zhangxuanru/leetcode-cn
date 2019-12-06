<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/12/5
 * Time: 21:28
 */
/**
给定一个字符串来代表一个学生的出勤记录，这个记录仅包含以下三个字符：

'A' : Absent，缺勤
'L' : Late，迟到
'P' : Present，到场
如果一个学生的出勤记录中不超过一个'A'(缺勤)并且不超过两个连续的'L'(迟到),那么这个学生会被奖赏。

你需要根据这个学生的出勤记录判断他是否会被奖赏。

示例 1:

输入: "PPALLP"
输出: True
示例 2:

输入: "PPALLL"
输出: False

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/student-attendance-record-i
 **/
class Solution {

    /**
     * @param String $s
     * @return Boolean
     */
    function checkRecord($s) {
            $i = 0;
            $len = strlen($s);
            $data = array();
           while($i < $len){
                  $c = $s[$i];
                  if( $c== 'A' && isset($data['A']) && $data['A'] > 0){
                      return false;
                  }
                if($c == 'L' && $s[$i+1] == 'L' && $s[$i+2] == 'L'){
                   return false;
                }
               if(isset($data[$c])){
                   $data[$c]++;
               }else{
                   $data[$c]=1;
               }
               $i++;
           }
        return true;
    }


    function checkRecord2($s)
    {
        $i =0;
        $len = strlen($s);
        $cA =0;
        $cL = 0;
        while($i < $len){
            $c = $s[$i];
            if($c == 'A'){
                $cA++;
            }
            if($c =='L'){
                $cL++;
            }else{
                $cL=0;
            }
           if($cA >1 || $cL>2){
               return false;
           }
            $i++;
        }
        return true;
    }
}

$class = new Solution();
$s = "PPALLP";
$s ="PPALLL";
$s="LPLPLPLPLPL";
$b = $class->checkRecord($s);
var_dump($b);
$b = $class->checkRecord2($s);
var_dump($b);