<?php
  //globals and includes

  include 'credentials.php';
  include 'Flight.php';

  $redBound = 0.2;
  $yellowBound = 0.35;
  $greenBound = 0.5;
  $hardCap = 150;
  $minTripLength = 2;
  $limitSize = 10;

?>
<?php
  //functions

  function getFlightsOfInterest($conn,$src,$dest,$bound,$previousPriceCap){

    global $minTripLength;
    global $hardCap;
    global $limitSize;
    $query = "SELECT * , (Price/(SELECT AveragePrice FROM Averages WHERE AirPort = '$dest')) FROM ${src}_${dest} WHERE Price < ($bound * (SELECT AveragePrice FROM Averages WHERE AirPort = '$dest')) AND DATEDIFF(ReturnDate,DepartDate) > $minTripLength AND Price < $hardCap AND Price > $previousPriceCap ORDER BY (Price/(SELECT AveragePrice FROM Averages WHERE AirPort = '$dest')) ASC limit $limitSize;";

    $pricesArray = arraySetup($conn,$query);

    return $pricesArray;

  }

  //this function will be used to setup my source and destination airport arrays
  function arraySetup($conn, $query){

    // echo $query . "\n";

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

  global $redBound;
  global $yellowBound;
  global $greenBound;

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

      array_merge($redAlerts, getFlightsOfInterest($conn,$src,$dest,$redBound,$hardCap,0));
      array_merge($yellowAlerts, getFlightsOfInterest($conn,$src,$dest,$yellowBound,$hardCap,$redBound));
      array_merge($greenAlerts, getFlightsOfInterest($conn,$src,$dest,$greenBound,$hardCap,$yellowBound));

    }
  }

  //at this point i'll have an array of red, yellow and green alerts
  // so i want to send out a total of 15 reports. first priority will go to the red red alerts, then the yellow..

  $reportsArray = array();

  while(count($reportsArray) < 15){

    $foreach($redAlerts as $redAlert){

    }

  }

  //close connection
  $conn->close();

?>
