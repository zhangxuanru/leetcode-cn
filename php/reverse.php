<?php
function Reverse($x)
{
  if($x < 0){
        return -Reverse(-$x);
    }
  $res = 0;
  while($x!=0){
     $res = $res*10 +$x%10;
     $x = (int)($x/10);
  }
  if($res > 2147483647 || $res < -2147483648){
      return 0;
  }
  return $res;
}
$x=-321;
echo Reverse($x);
