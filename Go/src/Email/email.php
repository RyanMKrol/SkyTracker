<?php
  //globals and includes

  include "./../Credentials/credentials.php";

?>
<?php
  //main

  $title = trim(file_get_contents($argv[1]));
  $body = $argv[2];
  $emailAddress = $argv[3];

  exec("mail -a \"From: SkyTracker <no-reply@skytracker.co>\" -a \"Content-type: text/html\" -s \"$title\" $emailAddress < $body");
?>
