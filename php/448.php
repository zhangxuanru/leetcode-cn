<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/11/28
 * Time: 16:51
 */

function FindDisappearedNumbers($nums){
    $data = array();
      for($i=0;$i<count($nums);$i++){
           $key = abs($nums[$i])-1;
          if($nums[$key] > 0){
              $nums[$key] = -$nums[$key];
          }
      }
    for($i=0; $i<count($nums);$i++){
        if($nums[$i] > 0){
            array_push($data,$i+1);
            echo $i+1;
            echo "\r\n";
        }
    }

    print_r($nums);
    print_r($data);
}

$arr = array(4, 3, 2, 7, 8, 2, 3, 1);
FindDisappearedNumbers($arr);