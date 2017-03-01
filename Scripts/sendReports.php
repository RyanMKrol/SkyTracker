<?php
  //globals and includes

  include 'credentials.php';
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

  // Create connection
  $conn = new mysqli($servername, $username, $password, $database);

  // Check connection
  if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
  }

  $emailsArr = arraySetup($conn, "SELECT * FROM Users;");

  //getting a minimum for each destination, from all of the sources
  foreach($emailsArr as $address) {

    echo "user\n";
    echo $address['UserEmailAddress'] . "\n";
    exec("echo \"This is the body of the email\" | mail -A ./confirmation.txt -s \"This is the subject line\" ryankrol@hotmail.co.uk");

    // exec("mysql -u root -p\"$password\" -f \"SkyTracker\" < /var/www/html/skytracker.co/Data/Averages.sql");
  }
  //close connection
  $conn->close();

?>
