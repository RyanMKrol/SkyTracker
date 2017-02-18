<?php
  //includes and globals

  include 'credentials.php';
?>
<?php
  //functions

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

  $myfile = fopen("/var/www/html/skytracker.co/Data/Averages.sql", "w+");

  fwrite($myfile, "DROP TABLE Averages;\n\n");
  fwrite($myfile, "CREATE TABLE Averages (\n");
  fwrite($myfile, "\tAirPortID int NOT NULL AUTO_INCREMENT,\n");
  fwrite($myfile, "\tAirPort varchar(255) NOT NULL,\n");
  fwrite($myfile, "\tAveragePrice int NOT NULL,\n");
  fwrite($myfile, "\tPRIMARY KEY(AirPortID)\n");
  fwrite($myfile, ");\n\n");

  foreach($destinationsArr as $destAirport) { //foreach element in $arr

    $dest = $destAirport["DestAirportCode"];
    $sum = 0;
    foreach($sourcesArr as $srcAirport) { //foreach element in $arr

      $src = $srcAirport["SrcAirportCode"];

      $query = "SELECT AVG(Price) FROM ${src}_${dest}";
      $avgPrice = arraySetup($conn,$query);
      $avgPrice = $avgPrice[0]["AVG(Price)"];

      $sum += $avgPrice;
    }

    $sum /= count($sourcesArr);

    fwrite($myfile, "INSERT INTO Averages (AirPort, AveragePrice) VALUES ('$dest', '$sum');\n");

  }
  
  fclose($myfile);
  $conn->close();

  exec("mysql -u root -p\"$password\" -f \"SkyTracker\" < /var/www/html/skytracker.co/Data/Averages.sql");

?>
