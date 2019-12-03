<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/12/3
 * Time: 11:03
 */
class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function longestPalindrome($s) {
          /**
           * 思路：
           *  相同的字母按左右拆分，
           *  如果不是相同的，向中间插入，
           *  记录字符串的长度， 找到 最长的那个数
           */
        $i = 0;
        $data = array();
        while(isset($s[$i])){
            $num = isset( $data[$s[$i]]) ?  $data[$s[$i]] : 0;
            $data[$s[$i]] = $num+1;
            $i++;
        }
        if(count($data) == 1){
            return strlen($s);
        }
        $len = 0;
        $isPri = 0;
        foreach($data as $char => $num){
             if($num%2 == 0){
                 $len+=$num;
                 continue;
             }
             if($num>2 && $num%2!=0){
                 $len+= $num-1;
                 $isPri = 1;
             }else{
                 $isPri =1;
             }
        }
        return $len+$isPri;
    }
}

$class = new Solution();
$s = "abccccdd";
$s = "ccc";
//$s = "a";
//$s ="abb";
$s = "bananas";
//$s = "ababababa";
$ret = $class->longestPalindrome($s);
echo "\r\n";
echo $ret;


