<?php

  include 'credentials.php';

  // Create connection
  $conn = new mysqli($servername, $username, $password, $database);

  // Check connection
  if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
  }

  $sqlSources = "SELECT * FROM SourceAirports;";
  $sources = $conn->query($sqlSources);

  $sqlDestinations = "SELECT * FROM DestinationAirports;";
  $destinations = $conn->query($sqlDestinations);

  $sourcesArr = array();
  $destinationsArr = array();


  while($row = mysqli_fetch_array($sources)){
    $sourcesArr[] = $row;
  }

  while($row = mysqli_fetch_array($destinations)){
    $destinationsArr[] = $row;
  }

  foreach($sourcesArr as $row){
      echo "id: " . $row["SrcAirportID"]. " - Name: " . $row["SrcAirportName"] . "\n";

      foreach($destinationsArr as $innerRow){
          echo "id: " . $innerRow["DestAirportID"]. " - Name: " . $innerRow["DestAirportName"] . "\n";
      }
  }

  $conn->close();
?>
