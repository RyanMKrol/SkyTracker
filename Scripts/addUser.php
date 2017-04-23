<?php
  //globals and includes
  include './../Go/src/Credentials/credentials.php';

?>
<?php
  //main

  $data = json_decode($_POST["_data"], true);

  $email    = $data['emailAddress'];
  $budget   = $data['budget'];
  $tripMin  = $data['tripMinLen'];
  $tripMax  = $data['tripMaxLen'];
  $months   = $data['months'];
  $airports = $data['airports'];
  $salt     = hash("sha256", $email . time());

  echo ("<p>$email</p>\n");
  echo ("<p>$budget</p>\n");
  echo ("<p>$tripMin</p>\n");
  echo ("<p>$tripMax</p>\n");
  echo ("<p>$months</p>");
  echo ("<p>$airports</p>");

  // Create connection
  $conn = new mysqli($server, $user, $password, $database);

  // Check connection
  if ($conn->connect_error) {
    die("Connection failed: " . $conn->connect_error);
  }

  $email    = $conn->real_escape_string($email);
  $budget   = $conn->real_escape_string($budget);
  $tripMin  = $conn->real_escape_string($tripMin);
  $tripMax  = $conn->real_escape_string($tripMax);
  $salt     = $conn->real_escape_string($salt);

  $sql = "INSERT INTO Users (UserEmailAddress, UserBudget, UserTripMin, UserTripMax, UserSalt) VALUES ('$email',$budget,$tripMin,$tripMax,'$salt');";

  if ($conn->query($sql) === TRUE) {
      echo "<p>New record created successfully</p>\n";
  } else {
      echo "\n" . $sql . "\n";
      echo "<p>oh dear</p>\n";
  }

  foreach($months as $month => $val) {
    if($val == true){
      $sql = "INSERT INTO UserTravelMonths (UserEmailAddress, TravelMonth) VALUES ('$email', $month);";
      if ($conn->query($sql) === TRUE) {
          echo "<p>successfullt added travel month</p>\n";
      } else {
          echo "\n" . $sql . "\n";
          echo "<p>failed to add travel month</p>\n";
      }
    }
  }

  foreach($airports as $airport => $val) {
    if($val == true){
      $sql = "INSERT INTO UserSourceAirports (UserEmailAddress, SourceAirportCode) VALUES ('$email', '$airport');";
      if ($conn->query($sql) === TRUE) {
          echo "<p>successfullt added airport</p>\n";
      } else {
          echo "<p>failed to add airport</p>\n";
      }
    }
  }

  $conn->close();

  //
  // //need to deal with cases here - we can do this later:
  // // 1. The record is created successfully, in which case we just welcome the user to the platform
  // // 2. The record already exists in the database, in which case we update the user's details, and give confirmation of that
  // // 3. The record was not created in the database, we report an error
  //
  //
  // //need to add response code here saying that the email has been added
  // echo "<p>$email has been added</p>";
  //

?>
