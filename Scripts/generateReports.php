<?php

  include 'credentials.php';
  $maxTripLength = 28;

  // Create connection
  $conn = new mysqli($servername, $username, $password, $database);

  // Check connection
  if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
  }

  //this is basic setup stuff for getting the source and destination airports
  {
    $sources = $conn->query("SELECT * FROM SourceAirports;");
    $destinations = $conn->query("SELECT * FROM DestinationAirports;");

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
  }

  foreach($sourcesArr as $srcAirport) {
    foreach($destinationsArr as $destAirport) {

      $src = $srcAirport["SrcAirportCode"];
      $dest = $destAirport["DestAirportCode"];

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
  }

  //close connection
  $conn->close();

?>
