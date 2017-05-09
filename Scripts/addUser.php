<?php
  //globals and includes
  include './../Go/src/Credentials/credentials.php';

?>
<?php
  //main

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


  $data           = json_decode($_POST["_data"], true);
  $authentication = $conn->real_escape_string($data['salt']);
  $email          = $conn->real_escape_string($data['emailAddress']);

  // if there's no attempt at authentication, we make a new user
  if(strcmp($authentication, '0') == 0){

    if(createUserData($data,$conn)){
      echo "successfully created the user";
    } else {
      echo "something failed when creating the user";
    }

  // otherwise we see if the authorisation checks out and then update
  } else {

    $sql    = "SELECT * FROM Users WHERE UserSalt = '$authentication';";
    $result = mysqli_query($conn, $sql);

    // if there are results for the salt, and the email address in the result is the same as the one passed in, we succeed
    if ((mysqli_num_rows($result) != 0) && (strcmp($email, mysqli_fetch_assoc($result)['UserEmailAddress']) == 0)){

      echo "salt and address are good\n";
      if(updateUserData($data,$conn)){
        echo "successfully updated everything\n";
      } else {
        echo "something failed when updating everything\n";
      }
    } else {
      echo "salt and address are not good\n";
    }
  }

  // close the connection to the database
  $conn->close();

?>
<?php

  // updates all of the users information in the database
  function updateUserData($data, $conn) {

    $email          = $conn->real_escape_string($data['emailAddress']);
    $budget         = $conn->real_escape_string($data['budget']);
    $tripMin        = $conn->real_escape_string($data['tripMinLen']);
    $tripMax        = $conn->real_escape_string($data['tripMaxLen']);
    $months         = $data['months'];
    $airports       = $data['airports'];

    $sql = "UPDATE Users SET UserBudget=$budget, UserTripMin=$tripMin, UserTripMax=$tripMax WHERE UserEmailAddress='$email';";

    if ($conn->query($sql) !== TRUE) {
      echo $sql . "\n";
        echo "Failed to update the user\n";
        mysqli_rollback($conn);
        return false;
    }

    $sql = "DELETE FROM UserTravelMonths WHERE UserEmailAddress = '$email';";

    if ($conn->query($sql) !== TRUE) {
        echo "Failed to delete user travel months\n";
        mysqli_rollback($conn);
        return false;
    }

    $sql = "DELETE FROM UserSourceAirports WHERE UserEmailAddress = '$email';";

    if ($conn->query($sql) !== TRUE) {
        echo "Failed to delete user source airports\n";
        mysqli_rollback($conn);
        return false;
    }

    foreach($months as $month => $val) {
      if($val == true){

        $month = $conn->real_escape_string($month);
        $sql = "INSERT INTO UserTravelMonths (UserEmailAddress, TravelMonth) VALUES ('$email', $month);";

        if ($conn->query($sql) !== TRUE) {
            echo "Failed to add a new month to the user\n";
            mysqli_rollback($conn);
            return false;
        }
      }
    }

    foreach($airports as $airport => $val) {
      if($val == true){

        $airport = $conn->real_escape_string($airport);
        echo $airport . "\n";
        echo $email . "\n";
        $sql = "INSERT INTO UserSourceAirports (UserEmailAddress, SourceAirportCode) VALUES ('$email', '$airport');";

        if ($conn->query($sql) !== TRUE) {
            echo "Failed to add a new airport to the user\n";
            mysqli_rollback($conn);
            return false;
        }
      }
    }

    // commit the data to the database
    mysqli_commit($conn);

    return true;
  }

  // creates a new user in the database
  function createUserData($data, $conn) {

    echo "entered create user \n";

    $email          = $conn->real_escape_string($data['emailAddress']);
    $budget         = $conn->real_escape_string($data['budget']);
    $tripMin        = $conn->real_escape_string($data['tripMinLen']);
    $tripMax        = $conn->real_escape_string($data['tripMaxLen']);
    $months         = $data['months'];
    $airports       = $data['airports'];
    $frequency      = $conn->real_escape_string($data['frequency']);

    // create a hash which will act as a users 'password'. Whenever they want to update their details,
    //  we'll send this hash and it'll be sent back to us as a token to prove it's from their email address. NOT
    //   as secure as a password, but it means that people can't just update other people's preferences just by knowing
    //    their email address
    $salt           = $conn->real_escape_string(hash("sha256", $email . time()));

    $sql = "INSERT INTO Users (UserEmailAddress, UserBudget, UserTripMin, UserTripMax, UserReportFrequency, UserSalt) VALUES ('$email',$budget,$tripMin,$tripMax,$frequency,'$salt');";

    echo $sql . "\n";

    if ($conn->query($sql) !== TRUE) {
        echo "Failed to add a new record\n";
        mysqli_rollback($conn);
        return false;
    }

    foreach($months as $month => $val) {
      if($val == true){

        $month = $conn->real_escape_string($month);
        $sql = "INSERT INTO UserTravelMonths (UserEmailAddress, TravelMonth) VALUES ('$email', $month);";

        if ($conn->query($sql) !== TRUE) {
            echo "Failed to add a new month to the user\n";
            mysqli_rollback($conn);
            return false;
        }
      }
    }

    foreach($airports as $airport => $val) {
      if($val == true){

        $airport = $conn->real_escape_string($airport);
        $sql = "INSERT INTO UserSourceAirports (UserEmailAddress, SourceAirportCode) VALUES ('$email', '$airport');";

        if ($conn->query($sql) !== TRUE) {
            echo "Failed to add a new airport to the user\n";
            mysqli_rollback($conn);
            return false;
        }
      }
    }

    // commit the data to the database
    mysqli_commit($conn);

    return true;
  }
?>
