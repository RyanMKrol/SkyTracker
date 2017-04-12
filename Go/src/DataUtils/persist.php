<?php

  include "./../Credentials/credentials.php"

  //creating "master" sql file
  exec("cat $raw_file_loc > $file_loc");

  //update updating the database
  exec("mysql -u $user -p\"$password\" -f \"$database\" < $file_loc");

?>
