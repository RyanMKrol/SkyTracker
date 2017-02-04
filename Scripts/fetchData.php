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

  foreach($sourcesArr as $srcAirport) { //foreach element in $arr
    foreach($destinationsArr as $destAirport) { //foreach element in $arr

      echo $srcAirport["SrcAirportCode"] . "\n";
      echo $destAirport["DestAirportCode"] . "\n";
      //this is going to have to be replaced very soon

      $departYear = 2017;
      $departMonth = "02";
      $returnYear = 2017;
      $returnMonth = "02";
      $src = $srcAirport["SrcAirportCode"];
      $dest = $destAirport["DestAirportCode"];


      //found in functions.php
      $mysqlFile = createMySQLFile($src, $dest);

      $data = getData($src, $dest, $departYear, $departMonth, $returnYear, $returnMonth);

      if(!is_null($data)){

          //write the data to the sql file
          writeData($data, $mysqlFile, $src, $dest, $departYear, $departMonth, $returnYear, $returnMonth);

          //update the database
          exec("mysql -u root -p\"$password\" -f \"SkyTracker\" < ${src}_${dest}.sql");

      } else {
          //have to decide on functionality for this later
      }

      //for padding months later: str_pad($input, 10, "-=", STR_PAD_LEFT);

      fclose($mysqlFile);


    }
  }
?>
