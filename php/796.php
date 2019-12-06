<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/12/6
 * Time: 15:15
 */
/**
给定两个字符串, A 和 B。

A 的旋转操作就是将 A 最左边的字符移动到最右边。 
 * 例如, 若 A = 'abcde'，在移动一次之后结果就是'bcdea' 。如果在若干次旋转操作之后，A 能变成B，那么返回True。

示例 1:
输入: A = 'abcde', B = 'cdeab'
输出: true

示例 2:
输入: A = 'abcde', B = 'abced'
输出: false
注意：

A 和 B 长度不超过 100。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-string
  **/
class Solution {

    /**
     * @param String $A
     * @param String $B
     * @return Boolean
     */
    function rotateString($A, $B) {
          if($A == $B){
              return true;
          }
          $len = strlen($A);
          $i = 0;
         while($i < $len-1){
               $c = $A[0];
               $A.= $c;
               $A = substr($A,1);
               if($A == $B){
                   return true;
                }
               $i++;
           }
        return false;
    }


    function rotateString2($A, $B) {
        if($A == $B){
            return true;
        }
        $len = strlen($A);
        $i = 0;
        while($i < $len-1){
            $A.=  $A[0];
            $A[0] = " ";
            $A= trim($A);
            if($A == $B){
                return true;
            }
            $i++;
        }
        return false;
    }


}

$class = new Solution();
$A = 'abcde';
$B = 'cdeab';
$r = $class->rotateString($A,$B);
var_dump($r);

$r = $class->rotateString2($A,$B);
var_dump($r);
