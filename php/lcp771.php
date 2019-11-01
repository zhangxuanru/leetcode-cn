<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/10/30
 * Time: 12:07
 */
/**
 * @param  string $J
 * @param string $S
 * @return integer
 */
function numJewelsInStones($J,$S){
 $jArr = str_split($J,1);
 $jArr = array_flip($jArr);
 $len = strlen($S);
 $r = 0;
 for($i=0;$i<$len;$i++){
     $key = $S[$i];
  if(isset($jArr[$key])){
      $r++;
  }
 }
  return $r;
}
function numJewelsInStones2($J,$S){
    return strlen($S) - strlen(str_replace(str_split($J,1),"",$S));
}
$j ="aA";
$s="aAAbbbb";
$r = numJewelsInStones($j,$s);
echo $r;
echo "\r\n";
echo numJewelsInStones2($j,$s);