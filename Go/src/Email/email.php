<?php
  //globals and includes

  include "./../Credentials/credentials.php";

?>
<?php
  //main

  $reportName = $argv[1];
  $title = trim(file_get_contents($argv[2]));
  $body = $argv[3];
  $emailAddress = $argv[4];

  exec("cat \"$body\" | mail -A $reportName -a \"From: SkyTracker <no-reply@skytracker.co>\" -s \"$title\" $emailAddress");

?>
