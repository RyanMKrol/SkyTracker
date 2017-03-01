<?php
  //globals and includes

  include 'credentials.php';
  $currentDate = date_create();

?>
<?php
  //functions

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
  
  global $currentDate;

  // Create connection
  $conn = new mysqli($servername, $username, $password, $database);

  // Check connection
  if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
  }

  $emailsArr = arraySetup($conn, "SELECT * FROM Users;");

  //getting a minimum for each destination, from all of the sources
  foreach($emailsArr as $address) {

    $reportName = "/var/www/html/skytracker.co/Reports/Report_" . date_format($currentDate,"d-m-Y") . ".csv";

    $emailAddress = $address['UserEmailAddress'];

    exec("echo \"Please find attached the report of cheap flights for Europe!\" | mail -A $reportName -s \"Your Daily Cheap Flights Report!\" $emailAddress");
  }
  //close connection
  $conn->close();

?>
