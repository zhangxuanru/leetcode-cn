<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/10/30
 * Time: 11:02
 */

/**
小A 和 小B 在玩猜数字。小B 每次从 1, 2, 3 中随机选择一个，小A 每次也从 1, 2, 3 中选择一个猜。他们一共进行三次这个游戏，请返回 小A 猜对了几次？
输入的guess数组为 小A 每次的猜测，answer数组为 小B 每次的选择。guess和answer的长度都等于3。
示例 1：
输入：guess = [1,2,3], answer = [1,2,3]
输出：3
解释：小A 每次都猜对了。
示例 2：
输入：guess = [2,2,3], answer = [3,2,1]
输出：1
解释：小A 只猜对了第二次
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/guess-numbers
*/
/**
 * @param array $game
 * @param array $answer
 * @return mixed
 */
function game($game=array(),$answer=array()){
    $r = 0;
    foreach($game as $key => $value){
        if($answer[$key] == $value){
            $r++;
        }
    }
    return $r;
}

$guess = array(1, 2, 3);
$answer = array(1, 2, 3);
$guess = array(2, 2, 3);
$answer = array(3, 2, 1);
$r = game($guess,$answer);
echo $r;

