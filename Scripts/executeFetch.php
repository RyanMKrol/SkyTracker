<?php
  //globals and includes
  include 'credentials.php';

?>
<?php
  //functions

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
  //main

  // Create connection
  $conn = new mysqli($servername, $username, $password, $database);

  // Check connection
  if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
  }

  $sourcesArr = arraySetup($conn, 'SourceAirports');
  $destinationsArr = arraySetup($conn, 'DestinationAirports');

  foreach($sourcesArr as $srcAirport) { //foreach element in $arr
    foreach($destinationsArr as $destAirport) { //foreach element in $arr

      $one = $srcAirport['SrcAirportCode'];
      $two = $destAirport['DestAirportCode'];

      print("thing happening");
      exec("php fetchData.php $one $two > /dev/null 2>&1 &");
    }
  }
?>
