<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/10/16
 * Time: 11:04
 */

/*
 * 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 *
 * */

/**
 * @param array $nums
 * @return array
 */
function threeSum($nums = array())
{
    sort($nums);
    $len = count($nums);
    $ret = array();
    for($i=0;$i<$len;$i++){
        for($j=$i+1;$j<$len;$j++){
            for($k=$j+1;$k<$len;$k++){
                $res = [];
                $flag = true;
                if($nums[$i] + $nums[$j] + $nums[$k] ==0 ){
                    $res = array($nums[$i],$nums[$j],$nums[$k]);
                }
                foreach($ret as $key => $val){
                    if(is_array($val)){
                        if($val[0] == $res[0] && $val[1] == $res[1] && $val[2] == $res[2]){
                            $flag = false;
                            break;
                        }
                    }
                }
                if($flag && !empty($res)){
                    $ret[] = $res;
                }
            }
        }
    }
    return $ret;
}

/**
 * 推荐这种方法。。。
 * @param array $nums
 * @return array
 */
function threeSum2($nums = array()){
    sort($nums);
    $len = count($nums);
    $result = array();
    for($i=0;$i<$len;$i++){
         if($i > 0 && $nums[$i-1] == $nums[$i]){
             continue;
         }
        $l = $i+1;
        $r = $len-1;
        while($l < $r){
             $res = $nums[$i] + $nums[$l] + $nums[$r];
            if($res == 0){
                $result[] = array($nums[$i],$nums[$l],$nums[$r]);
                $l++;
                $r--;
                while($l < $r && $nums[$l] == $nums[$l-1]){
                    $l++;
                }
                while($l < $r && $nums[$r] == $nums[$r+1]){
                    $r--;
                }
            }
            if($res > 0){
                $r--;
            }
            if($res<0){
                $l++;
            }
        }
    }
    return $result;
}
$arr = [-1, 0, 1, 2, -1, -4];
$ret = threeSum($arr);
print_r($ret);

echo "----------------------------\r\n";
$val = threeSum2($arr);
print_r($val);