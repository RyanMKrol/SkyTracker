<?php

  include 'credentials.php';
  include 'httpCodes.php';

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

  $departYear = 2017;
  $departMonth = "02";
  $returnYear = 2017;
  $returnMonth = "02";
  $srcAirport = $sourcesArr[0]["SrcAirportCode"];
  $destAirport = $destinationsArr[0]["DestAirportCode"];

  //for padding months later: str_pad($input, 10, "-=", STR_PAD_LEFT);

  $call = "http://partners.api.skyscanner.net/apiservices/browsegrid/v1.0/GB/GBP/en-GB/$srcAirport/$destAirport/$departYear-$departMonth/$returnYear-$returnMonth?apiKey=$apikey";

  //looping through to show how i'll get the pairings.
  foreach($sourcesArr as $row){
      // echo "id: " . $row["SrcAirportID"]. " - Name: " . $row["SrcAirportName"] . "\n";

      foreach($destinationsArr as $innerRow){
          // echo "id: " . $innerRow["DestAirportID"]. " - Name: " . $innerRow["DestAirportName"] . "\n";

      }
  }

  // initialist the api request
  $curl = curl_init($call);

  // returns the api request as a string
  curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);

  // execute the api request
  $curl_response = curl_exec($curl);

  echo $call . "\n";

  switch ($http_code = curl_getinfo($curl, CURLINFO_HTTP_CODE)) {
    case $httpSuccess:  # All's fine

      $data = json_decode($curl_response,true);
      echo sizeof($data);
      foreach($data as $key => $val) { //foreach element in $arr
        echo $key . "\n";
      }
      echo "all good\n";
      break;
    case $httpExcess:  # Using too much
      echo "all good\n";
      break;
    default:
      echo 'Unexpected HTTP code: ', $http_code, "\n";
  }

  $conn->close();

?>
