<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/12/5
 * Time: 15:45
 */

/**
 * 给定一个字符串和一个整数 k，你需要对从字符串开头算起的每个 2k 个字符的前k个字符进行反转。
 *
 如果剩余少于 k 个字符，则将剩余的所有全部反转。如果有小于 2k 但大于或等于 k 个字符，则反转前 k 个字符，并将剩余的字符保持原样。
示例:

输入: s = "abcdefg", k = 2
输出: "bacdfeg"
要求:

该字符串只包含小写的英文字母。
给定字符串的长度和 k 在[1, 10000]范围内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-string-ii
 **/
class Solution {

    /**
     * todo 思路：按2K个字符拆分，然后前k个字符反转，最后判断最后一组字符是否小于K,再具体决定是否反转
     * @param String $s
     * @param Integer $k
     * @return String
     */
    //todo error
    function reverseStr_test($s, $k) {
         $sArr = str_split($s,2*$k);
         $sCount = count($sArr)-1;
        foreach($sArr as $key => $value){
            if($key <= $sCount ||($key ==$sCount &&  strlen($value) < 2*$k && strlen($value)>=$k)){
                  $j = $k-1;
                  $start = 0;
                echo "----";
                 while($j>0){
                   $tmp = $value[$j];
                   $value[$j] = $value[$start];
                   $value[$start] = $tmp;
                   $j--;
                   $start++;
                 }
            }else{
                //todo 说明是最后一组字符串， 判断是否需要反转
                if(strlen($value) < $k){
                        $start = 0;
                        $j = strlen($value)-1;
                        while($j>0){
                            $tmp = $value[$start];
                            $value[$start] = $value[$j];
                            $value[$j] = $tmp;
                            $j--;
                            $start++;
                        }
                }
            }
            $sArr[$key] = $value;
        }
       return implode("",$sArr);
    }


    //todo success
    function reverseStr($s,$k)
    {
       $sLen = strlen($s);
      if ($k > $sLen){
          $k = $sLen;
      }
        $twoK = 2*$k;
       $i = 0;
       $isEnd = false;
       while($i < $sLen){
           $n = intval($k/2);
           $start = $i;
           $end  = $i+$k-1;
           if($end >= $sLen){
               $end = $sLen-1;
           }
           echo "end:",$end."n=",$n."\r\n";
           while($n > 0){
               $tmp = $s[$start];
               $s[$start] = $s[$end];
               $s[$end] = $tmp;
               $start++;
               $end--;
               $n--;
           }
           $i+=$twoK;
           $pow = $sLen-$i;
           if($pow < $k){
               $start = $i;
               $end = $sLen-1;
               $isEnd = true;
           }
           if($pow < $twoK && $pow>=$k){
               $start = $i;
               $end = $i+$k-1;
               $isEnd = true;
           }
           while($start<$end){
               $tmp = $s[$start];
               $s[$start] = $s[$end];
               $s[$end] = $tmp;
               $start++;
               $end--;
           }
           if($isEnd){
               break;
           }
       }
        return $s;
    }

    //todo
    function reverseStr3($s,$k)
    {
        $len = strlen($s);
        for($i=0;$i<$len;$i+=2*$k){
             if($i +$k <= $len){
                $this->reverse($s,$i,$i+$k-1);
             }else{
                $this->reverse($s,$i,$len-1);
             }
        }
        return $s;
    }

    private function reverse(&$s,$start,$end){
        echo $start."----end:".$end."\r\n";
            while($start<$end){
                $tmp = $s[$start];
                $s[$start] = $s[$end];
                $s[$end] = $tmp;
                $start++;
                $end--;
            }
    }
}


$class = new Solution();
$s = "abcdefg";  //输出: "bacdfeg"
$k = 2;

//$s = "abcd";
//$k =4;
//$s = "a";
//$k=2;
//$s = "abcdefg"; //"gfedcba"
//$k = 9;
//$s = "abcdefg"; //"gfedcba"
//$k = 1213;
$news = $class->reverseStr($s,$k);
echo $news;
echo "\r\n";
$news = $class->reverseStr3($s,$k);

echo $news;