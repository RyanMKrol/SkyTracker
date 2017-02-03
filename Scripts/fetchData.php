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

  //this is going to have to be replaced very soon
  $departYear = 2017;
  $departMonth = "02";
  $returnYear = 2017;
  $returnMonth = "02";
  $srcAirport = $sourcesArr[0]["SrcAirportCode"];
  $destAirport = $destinationsArr[0]["DestAirportCode"];

  getData($srcAirport, $destAirport, $departYear, $departMonth, $returnYear, $returnMonth);
  //for padding months later: str_pad($input, 10, "-=", STR_PAD_LEFT);

?>
