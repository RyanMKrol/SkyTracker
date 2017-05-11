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

  $data  = json_decode($_POST["_data"], true);
  $salt  = $conn->real_escape_string($data["token"]);
  $email = $conn->real_escape_string($data['email']);

  $returnData = removeUser($conn, $email, $salt);

  print_r(json_encode($returnData));

  // close the connection to the database
  $conn->close();

?>

<?php

  //fucntion to remove the user from the database, given a token
  function removeUser($conn, $email, $salt) {

    $returnData->success = false;
    $returnData->message = "Something went wrong when unsubscribing, sorry! Please contact root@skytracker.co for assisstance and we'll get right on it";

    if(isset($salt) && isset($email)){

      $sql    = "SELECT * FROM Users WHERE UserSalt = '$salt' AND UserEmailAddress = '$email';";
      $result = mysqli_query($conn, $sql);

      // if there are results for the salt, and the email address in the result is the same as the one passed in, we succeed
      if (mysqli_num_rows($result) != 0){

        $sql   = "DELETE FROM UserTravelMonths WHERE UserEmailAddress = '$email'";
        if ($conn->query($sql) !== TRUE) {
            mysqli_rollback($conn);
            return $returnData;
        }

        $sql   = "DELETE FROM UserSourceAirports WHERE UserEmailAddress = '$email'";
        if ($conn->query($sql) !== TRUE) {
            mysqli_rollback($conn);
            return $returnData;
        }

        $sql   = "DELETE FROM Users WHERE UserEmailAddress = '$email'";
        if ($conn->query($sql) !== TRUE) {
            mysqli_rollback($conn);
            return $returnData;
        }

        mysqli_commit($conn);

      } else {
        return $returnData;
      }
    } else {
      return $returnData;
    }

    $returnData->success = true;
    $returnData->message = "Success! You have been unsubscribed from our mailing list, hopefully you found a great flight!";

    return $returnData;
  }

?>
