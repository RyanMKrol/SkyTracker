<?php

  include 'credentials.php';

  $conn = new mysqli($servername, $username, $password, $database);

  // Check connection
  if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
  }

  $result = $conn->query("select * from Averages;");

  while($row = mysqli_fetch_array($result)){
  	$price = $row['AveragePrice'];
  	$airport = $row['AirPort'];
  	echo "$airport $price\n";
  }

  //close connection
  $conn->close();

?>
