<?php
  //globals and includes
  include './../Go/src/Credentials/credentials.php';

?>
<?php
  //main

  $data = json_decode($_POST["_data"], true);

  $email          = $data['emailAddress'];
  $budget         = $data['budget'];
  $tripMin        = $data['tripMinLen'];
  $tripMax        = $data['tripMaxLen'];
  $months         = $data['months'];
  $airports       = $data['airports'];
  $authentication = $data['salt'];

  if($authentication == '0'){
    echo "we'll create something new";
  } else {
    echo "we'll try and update and see what happens";
  }

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

  // stop this connection from committing unless we tell it to.
  mysqli_autocommit($conn, FALSE);

  // commit here before anything happens. When we rollback, nothing will have changed in the database
  mysqli_commit($conn);

  $email    = $conn->real_escape_string($email);
  $budget   = $conn->real_escape_string($budget);
  $tripMin  = $conn->real_escape_string($tripMin);
  $tripMax  = $conn->real_escape_string($tripMax);
  $salt     = $conn->real_escape_string($salt);

  $sql = "INSERT INTO Users (UserEmailAddress, UserBudget, UserTripMin, UserTripMax, UserSalt) VALUES ('$email',$budget,$tripMin,$tripMax,'$salt');";

  if ($conn->query($sql) === TRUE) {
      echo "<p>New record created successfully</p>\n";
  } else {
      echo "Failed to add a new record";
      mysqli_rollback($conn);
      return;
  }

  foreach($months as $month => $val) {
    if($val == true){
      $sql = "INSERT INTO UserTravelMonths (UserEmailAddress, TravelMonth) VALUES ('$email', $month);";
      if ($conn->query($sql) === TRUE) {
          echo "<p>successfully added travel month</p>\n";
      } else {
          echo "Failed to add a new month to the user";
          mysqli_rollback($conn);
          return;
      }
    }
  }

  foreach($airports as $airport => $val) {
    if($val == true){
      $sql = "INSERT INTO UserSourceAirports (UserEmailAddress, SourceAirportCode) VALUES ('$email', '$airport');";
      if ($conn->query($sql) === TRUE) {
          echo "<p>successfully added airport</p>\n";
      } else {
          echo "Failed to add a new airport to the user";
          mysqli_rollback($conn);
          return;
      }
    }
  }

  // commit the results to the database
  mysqli_commit($conn);

  // close the connection to the database
  $conn->close();

?>
