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

  $reportName = $argv[1];
  $title = $argv[2];
  $body = $argv[3];

  //getting a minimum for each destination, from all of the sources
  foreach($emailsArr as $address) {

    $emailAddress = $address['UserEmailAddress'];

    exec("echo \"$body\" | mail -A $reportName -s \"$title\" $emailAddress");
  }

  //close connection
  $conn->close();

?>
