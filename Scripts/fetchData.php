<?php

  include 'credentials.php';

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

  //looping through to show how i'll get the pairings.
  foreach($sourcesArr as $row){
      echo "id: " . $row["SrcAirportID"]. " - Name: " . $row["SrcAirportName"] . "\n";

      foreach($destinationsArr as $innerRow){
          echo "id: " . $innerRow["DestAirportID"]. " - Name: " . $innerRow["DestAirportName"] . "\n";
      }
  }

  $conn->close();
?>
