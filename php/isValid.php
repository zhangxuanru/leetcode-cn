<?php
/*
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
 * */

/**
 *
 * 有效的括号
 * @link https://leetcode-cn.com/problems/valid-parentheses/%E3%80%82/
 * @param $str
 * @return boolean
 */
function isValid($str)
{
    $replaceArr = array("()","{}","[]","{}");
    while(strlen($str) >0){
      $len = strlen($str);
      $str = str_replace($replaceArr,"",$str);
      if($len == strlen($str)){
          break;
      }
  }
    return strlen($str) == 0;
}

var_dump(isValid("()[]{}"));
var_dump(isValid("([)]"));
var_dump(isValid("{[]}"));
