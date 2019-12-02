<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/12/2
 * Time: 20:48
 */
class Solution {

    /**
     * @param Integer[] $numbers
     * @param Integer $target
     * @return Integer[]
     */
    /*
    function twoSum_test($numbers, $target) {
        $data = array();
         foreach($numbers as $key => $val){
             $data[$val] = $key;
        }
        foreach($numbers as $key => $val){
               $diff = $target - $val;
            if(isset($data[$diff]) && $key < $data[$diff]){
                return array($key+1,$data[$diff]+1);
            }
        }
        return [];
    }
*/
    function twoSum($numbers, $target) {
        $data = array();
        foreach($numbers as $key => $val){
            $diffKey = $target - $val;
            $index = 0;
            if(isset($data[$diffKey])){
                  $index = $data[$diffKey];
            }
            if($index > 0){
                $nKey = $key+1;
                $index1 = $index > $nKey ? $nKey :$index;
                $index2 = $index > $nKey ? $index :$nKey;
                return array($index1,$index2);
            }
            $data[$val] = $key+1;
        }
        return [];
    }
}

$class = new Solution();
$numbers = [2, 7, 11, 15];
$target = 9;
$arr = $class->twoSum($numbers,$target);
print_r($arr);