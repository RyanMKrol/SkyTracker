<?php

  include 'credentials.php';
  include 'httpCodes.php';
  include 'functions.php';

  // Create connection
  $conn = new mysqli($servername, $username, $password, $database);

  // Check connection
  if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
  }

  //forming Select statement for sources
  $sqlSources = "SELECT * FROM SourceAirports;";
  $sources = $conn->query($sqlSources);

  //forming Select statement for destinations
  $sqlDestinations = "SELECT * FROM DestinationAirports;";
  $destinations = $conn->query($sqlDestinations);

  //close connection
  $conn->close();

  //where i'll store the data
  $sourcesArr = array();
  $destinationsArr = array();

  //storing the query results in more permanent storage
  while($row = mysqli_fetch_array($sources)){
    $sourcesArr[] = $row;
  }

  while($row = mysqli_fetch_array($destinations)){
    $destinationsArr[] = $row;
  }

  //creates date from the current date
  $srcDate = date_create();
  $destDate = date_create();

  $monthsNeeded = 2;

  if($srcDate->format('d') >= 14){
    $monthsNeeded = 3;
  }

  foreach($sourcesArr as $srcAirport) { //foreach element in $arr
    foreach($destinationsArr as $destAirport) { //foreach element in $arr

      $src = $srcAirport["SrcAirportCode"];
      $dest = $destAirport["DestAirportCode"];

      //found in functions.php
      $mysqlFile = createMySQLFile($src, $dest);

      //we can depart within the next 6 months
      foreach(range(1,6) as $j){

        //and we can come back within #monthsNeeded months of the set-off date
        foreach(range(1,$monthsNeeded) as $i){

          $departYear = $srcDate->format('Y');
          $departMonth = $srcDate->format('m');
          $returnYear = $destDate->format('Y');
          $returnMonth = $destDate->format('m');


          $data = getData($src, $dest, $departYear, $departMonth, $returnYear, $returnMonth);

          if(!is_null($data)){

              //write the data to the sql file
              writeData($data, $mysqlFile, $src, $dest, $departYear, $departMonth, $returnYear, $returnMonth);

          } else {
              //have to decide on functionality for this later
          }

          date_add($destDate,date_interval_create_from_date_string("1 month"));

        }
        date_add($srcDate,date_interval_create_from_date_string("1 month"));

        $destDate = date_create();
        date_add($destDate,date_interval_create_from_date_string("$j months"));
      }

      $srcDate = date_create();
      $destDate = date_create();

      //update the database
      exec("mysql -u root -p\"$password\" -f \"SkyTracker\" < ${src}_${dest}.sql");

      //close file now we're done with the src-dest pair
      fclose($mysqlFile);
    }
  }
?>
