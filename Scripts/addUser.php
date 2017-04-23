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

  // create a hash which will act as a users 'password'. Whenever they want to update their details,
  //  we'll send this hash and it'll be sent back to us as a token to prove it's from their email address. NOT
  //   as secure as a password, but it means that people can't just update other people's preferences just by knowing
  //    their email address
  $salt     = hash("sha256", $email . time());

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

?>
