<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/12/2
 * Time: 20:34
 */

class Solution {

    /**
     * @param String $date
     * @return Integer
     */
    function dayOfYear($date) {
       echo date("z",strtotime($date));
    }
}

$class = new Solution();
$date = "2003-03-01";
$date="2019-01-09";
$class->dayOfYear($date);