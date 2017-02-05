<?php

  //globals and includes

  include 'credentials.php';
  include 'Flight.php';

  $redBound = 0.2;
  $yellowBound = 0.35;
  $greenBound = 0.5;
  $hardCap = 150;
  $minTripLength = 2;

?>
<?php

  //in this section I'm going to store any and all functions that I want to use, I'm keeping
  // them separate so that I can maintain everything easily

  function getFlightsOfInterest($conn,$src,$dest,$bound,$previousPriceCap){

    $query = "SELECT * FROM ${src}_${dest} WHERE Price < ($bound * (SELECT AVG(Price) FROM ${src}_${dest})) AND DATEDIFF(ReturnDate,DepartDate) > $minTripLength AND Price < $hardCap AND Price > $previousPriceCap;";

    $pricesArray = arraySetup($conn,$query);

    return $pricesArray;

  }

  //this function will be used to setup my source and destination airport arrays
  function arraySetup($conn, $query){

    $result = $conn->query($query);

    $resultArray = array();

    while($row = mysqli_fetch_array($result)){
      $resultArray[] = $row;
    }

    return $resultArray;
  }
?>
<?php

  //main

  // Create connection
  $conn = new mysqli($servername, $username, $password, $database);

  // Check connection
  if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
  }

  $sourcesArr = arraySetup($conn, "SELECT * FROM SourceAirports;");
  $destinationsArr = arraySetup($conn, "SELECT * FROM DestinationAirports;");

  $redAlerts = array();
  $yellowAlerts = array();
  $greenAlerts = array();

  foreach($sourcesArr as $srcAirport) {
    foreach($destinationsArr as $destAirport) {

      $src = $srcAirport["SrcAirportCode"];
      $dest = $destAirport["DestAirportCode"];

      array_merge($redAlerts, getFlightsOfInterest($src,$dest,$redBound,$hardCap,0));
      array_merge($yellowAlerts, getFlightsOfInterest($src,$dest,$yellowBound,$hardCap,$redBound));
      array_merge($greenAlerts, getFlightsOfInterest($src,$dest,$greenBound,$hardCap,$yellowBound));

    }
  }

  //close connection
  $conn->close();

?>
