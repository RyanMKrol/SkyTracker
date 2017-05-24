<?php
  //globals and includes
  include './../Go/src/Credentials/credentials.php';

  $MAX_ACCOUNTS_PER_IP = 50;
?>
<?php
  //main

  $returnData->success = true;
  $returnData->message = "";

  // Create connection
  $conn = new mysqli($server, $user, $password, $database);

  // echo "connection fine";

  // Check connection
  if ($conn->connect_error) {
    die("Connection failed: " . $conn->connect_error);
  }

  // echo "connection still fine";

  // stop this connection from committing unless we tell it to.
  mysqli_autocommit($conn, FALSE);
  // commit here before anything happens. When we rollback, nothing will have changed in the database
  mysqli_commit($conn);

  $data           = json_decode($_POST["_data"], true);
  $authentication = $conn->real_escape_string($data['salt']);
  $email          = $conn->real_escape_string($data['emailAddress']);
  $captchaToken   = $data["captcha"];

  if(checkCaptcha($captchaToken)){
    // echo "captcha and IP fine";

    if(!userExists($conn, $email)){
      // echo "user does not exist so create one";
      //no user exists so we create one

      if(createUserData($data,$conn)){
        $returnData->message = "We've managed to add you to our mailing list! You'll get an email from us shortly :)";
        exec("php newUser.php $email");
      } else {
        $returnData->success = false;
        $returnData->message = "Unfortunately we couldn't add you to our mailing list, please try again later";
        //echo "something failed when creating the user";
      }
    } else if(strcmp($authentication, '0') !== 0){
      //user exists and some authentication is attempted

      $sql    = "SELECT * FROM Users WHERE UserSalt = '$authentication' AND UserEmailAddress = '$email';";
      $result = mysqli_query($conn, $sql);

      // if there are results for the salt, and the email address in the result is the same as the one passed in, we succeed
      if (mysqli_num_rows($result) !== 0){

        //echo "salt and address are good\n";
        if(updateUserData($data,$conn)){
          $returnData->message = "Your preferences have been updated!";
          //echo "successfully updated everything\n";
        } else {
          $returnData->success = false;
          $returnData->message = "Unfortunately we had some trouble updating your preferences, please try again later, sorry!";
          //echo "something failed when updating everything\n";
        }
      } else {
        $returnData->success = false;
        $returnData->message = "To update your preferences, please follow the link from one of the emails we send you! (This is to stop anybody changing your preferences!)";
        echo "salt not good\n";
      }
    } else {
      $returnData->success = false;
      $returnData->message = "To update your preferences, please follow the link from one of the emails we send you! (This is to stop anybody changing your preferences!)";
    }
  }

  print_r(json_encode($returnData));

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
        //echo $sql . "\n";
        //echo "Failed to update the user\n";
        mysqli_rollback($conn);
        return false;
    }

    $sql = "DELETE FROM UserTravelMonths WHERE UserEmailAddress = '$email';";

    if ($conn->query($sql) !== TRUE) {
        //echo "Failed to delete user travel months\n";
        mysqli_rollback($conn);
        return false;
    }

    $sql = "DELETE FROM UserSourceAirports WHERE UserEmailAddress = '$email';";

    if ($conn->query($sql) !== TRUE) {
        //echo "Failed to delete user source airports\n";
        mysqli_rollback($conn);
        return false;
    }

    foreach($months as $month => $val) {
      if($val == true){

        $month = $conn->real_escape_string($month);
        $sql = "INSERT INTO UserTravelMonths (UserEmailAddress, TravelMonth) VALUES ('$email', $month);";

        if ($conn->query($sql) !== TRUE) {
            //echo "Failed to add a new month to the user\n";
            mysqli_rollback($conn);
            return false;
        }
      }
    }

    foreach($airports as $airport => $val) {
      if($val == true){

        $airport = $conn->real_escape_string($airport);
        //echo $airport . "\n";
        //echo $email . "\n";
        $sql = "INSERT INTO UserSourceAirports (UserEmailAddress, SourceAirportCode) VALUES ('$email', '$airport');";

        if ($conn->query($sql) !== TRUE) {
            //echo "Failed to add a new airport to the user\n";
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

    //echo "entered create user \n";

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

    //echo $sql . "\n";

    if ($conn->query($sql) !== TRUE) {
        //echo "Failed to add a new record\n";
        mysqli_rollback($conn);
        return false;
    }

    foreach($months as $month => $val) {
      if($val == true){

        $month = $conn->real_escape_string($month);
        $sql = "INSERT INTO UserTravelMonths (UserEmailAddress, TravelMonth) VALUES ('$email', $month);";

        if ($conn->query($sql) !== TRUE) {
            //echo "Failed to add a new month to the user\n";
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
            //echo "Failed to add a new airport to the user\n";
            mysqli_rollback($conn);
            return false;
        }
      }
    }

    // commit the data to the database
    mysqli_commit($conn);

    return true;
  }


  //function to check the IP address isn't overloading the server
  function userExists($conn, $emailAddress){

    $sql    = "SELECT * FROM Users WHERE UserEmailAddress = '$emailAddress';";
    $result = mysqli_query($conn, $sql);

    // if there are results for the salt, and the email address in the result is the same as the one passed in, we succeed
    return mysqli_num_rows($result) !== 0;

  }

  // curl functionality pulled from https://www.madebymagnitude.com/blog/sending-post-data-from-php/
  function checkCaptcha($captchaToken){

    // echo "checkign captcha";

    global $secretCaptcha;

    # Our new data
    $data = array(
      'secret' => $secretCaptcha,
      'response' => $captchaToken
      );

    # Create a connection
    $url = 'https://www.google.com/recaptcha/api/siteverify';
    $ch = curl_init($url);

    # Form data string
    $postString = http_build_query($data, '', '&');

    # Setting our options
    curl_setopt($ch, CURLOPT_POST, 1);
    curl_setopt($ch, CURLOPT_POSTFIELDS, $postString);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

    # Get the response
    $response = json_decode(curl_exec($ch), true);

    if(!$response['success']){
      //echo "Sorry, your CAPTCHA token is invalid\n";
    }

    $success = $response['success'] === true;

    // echo "\nreturning: $success\n";

    return $success;
  }


?>
