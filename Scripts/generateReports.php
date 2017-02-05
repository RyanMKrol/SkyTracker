<?php

  // i'm using this space to store my global variables. I'm not sure what the official way of doing this in php is,
  // but this will work well for me organising everything

  $redBound = 0.2;
  $yellowBound = 0.35;
  $greenBound = 0.5;
  $hardCap = 150;

?>
<?php

  //in this section I'm going to store any and all functions that I want to use, I'm keeping
  // them separate so that I can maintain everything easily

  function getFlightsOfInterest($src,$dest,$bound){

    //PUT IN A CONDITION THAT MAKES THE DATE LATER THAN TODAY'S
    $query = "SELECT * FROM ${src}_${dest} WHERE (Price < (0.3 * (SELECT AVG(Price) FROM ${src}_${dest})) AND DATEDIFF(ReturnDate,DepartDate) > 2 AND Price < 150) OR (Price < 40);";

    $price = $conn->query($query);

    $average = "SELECT AVG(Price) FROM ${src}_${dest};";
    $avPrice = $conn->query($average);

    $thing = mysqli_fetch_array($avPrice);
    echo "THE AVERAGE PRICE FOR ${src}_${dest} IS " . $thing['AVG(Price)'] . "\n";

    while($row = mysqli_fetch_array($price)){
      echo "Going from " . $row['SourcePort'] . " to " . $row['DestPort'] . " on " . $row['DepartDate'] . " coming back on " . $row['ReturnDate'] . " for the low low price of: " . $row['Price'] . "\n";
    }

  }

  //this function will be used to setup my source and destination airport arrays
  function arraySetup($conn, $table){

    $result = $conn->query("SELECT * FROM $table;");

    $resultArray = array();

    while($row = mysqli_fetch_array($result)){
      $resultArray[] = $row;
    }

    return $resultArray;
  }
?>
<?php

  //the bulk of my script will go here!

  include 'credentials.php';
  include 'Flight.php';

  // Create connection
  $conn = new mysqli($servername, $username, $password, $database);

  // Check connection
  if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
  }

  $sourcesArr = arraySetup($conn, 'SourceAirports');
  $destinationsArr = arraySetup($conn, 'DestinationAirports');

  $redAlerts = array();
  $yellowAlerts = array();
  $greenAlerts = array();

  foreach($sourcesArr as $srcAirport) {
    foreach($destinationsArr as $destAirport) {

      $src = $srcAirport["SrcAirportCode"];
      $dest = $destAirport["DestAirportCode"];

      array_merge($redAlerts, getFlightsOfInterest($src,$dest,$redBound, $hardCap));
      array_merge($yellowAlerts, getFlightsOfInterest($src,$dest,$yellowBound, $hardCap));
      array_merge($greenAlerts, getFlightsOfInterest($src,$dest,$greenBound, $hardCap));

    }
  }

  //close connection
  $conn->close();

?>
