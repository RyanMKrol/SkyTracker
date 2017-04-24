<?php
  //globals and includes
  include './Go/src/Credentials/credentials.php';

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

  $salt  = $conn->real_escape_string($_GET["token"]);
  $email = $conn->real_escape_string($_GET['email']);

  if(isset($salt) && isset($email)){

    $sql    = "SELECT * FROM Users WHERE UserSalt = '$salt';";
    $result = mysqli_query($conn, $sql);

    // if there are results for the salt, and the email address in the result is the same as the one passed in, we succeed
    if ((mysqli_num_rows($result) != 0) && (strcmp($email, mysqli_fetch_assoc($result)['UserEmailAddress']) == 0)){

      $sql   = "DELETE FROM UserTravelMonths WHERE UserEmailAddress = '$email'";
      if ($conn->query($sql) !== TRUE) {
          echo "Failed to delete the travel months\n";
          mysqli_rollback($conn);
          return false;
      }

      $sql   = "DELETE FROM UserSourceAirports WHERE UserEmailAddress = '$email'";
      if ($conn->query($sql) !== TRUE) {
          echo "Failed to delete the source airports\n";
          mysqli_rollback($conn);
          return false;
      }

      $sql   = "DELETE FROM Users WHERE UserEmailAddress = '$email'";
      if ($conn->query($sql) !== TRUE) {
          echo "Failed to delete the user\n";
          mysqli_rollback($conn);
          return false;
      }

      echo "success!";

      mysqli_commit($conn);

    } else {
      echo "salt and address are present, but not good\n";
    }
  } else {
    echo "either salt or email is not present\n";
  }

  // close the connection to the database
  $conn->close();

?>
