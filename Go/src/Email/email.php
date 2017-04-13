<?php
  //globals and includes

  include "./../Credentials/credentials.php";

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
  $conn = new mysqli($server, $user, $password, $database);

  // Check connection
  if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
  }

  $emailsArr = arraySetup($conn, "SELECT * FROM Users;");


  //getting a minimum for each destination, from all of the sources
  foreach($emailsArr as $address) {


    $reportName = $argv[1];

    $emailAddress = $address['UserEmailAddress'];

    exec("echo \"Please find attached the report of cheap flights for Europe!\" | mail -A $reportName -s \"Your Daily Cheap Flights Report!\" $emailAddress");
  }

  //close connection
  $conn->close();

?>
