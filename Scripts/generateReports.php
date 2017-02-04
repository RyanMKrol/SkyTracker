<?php

  include 'credentials.php';
  $maxTripLength = 28;

  // Create connection
  $conn = new mysqli($servername, $username, $password, $database);

  // Check connection
  if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
  }

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

  foreach(range(1,$maxTripLength) as $i){

    $users = $conn->query("SELECT * FROM Users WHERE UserTripLength = $i;");

    $numUsers = mysqli_num_rows($users);

    $usersArray = array();

    if($numUsers > 0){

      while($row = mysqli_fetch_array($users)){
        $usersArray[] = $row;
      }

      foreach($sourcesArr as $srcAirport) { //foreach element in $arr
        foreach($destinationsArr as $destAirport) { //foreach element in $arr

          $src = $srcAirport["SrcAirportCode"];
          $dest = $destAirport["DestAirportCode"];

          $query = "SELECT * FROM ${src}_${dest} WHERE Price = (SELECT MIN(Price) FROM ${src}_${dest} WHERE DATEDIFF(ReturnDate,DepartDate)=$i) AND DATEDIFF(ReturnDate,DepartDate)=$i;";
          // $conn->query($query);

          echo "$query\n";

        }
      }


    }

  }

  //close connection
  $conn->close();

?>
