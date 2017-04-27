<?php

  include "./../Credentials/credentials.php";


  $files = scandir($raw_file_dir);

  foreach($files as $file) {

      $fullFilePath = $raw_file_dir . $file;
      exec("mysql -u $user -p\"$password\" -f \"$database\" < $fullFilePath");
      echo $raw_file_dir . $file . "\n";

  }

?>
