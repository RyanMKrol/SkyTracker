<?php

  include 'credentials.php';

  $email = $_POST['addr'];
  $tripLength = $_POST['tripLength'];
  $budget = $_POST['budget'];

  // Create connection
  $conn = new mysqli($servername, $username, $password, $database);

  // Check connection
  if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
  }

  $email = $conn->real_escape_string($email);

  $sql = "INSERT INTO Users (UserEmailAddress, UserBudget, UserTripLength) VALUES ('$email', '$budget', '$tripLength');";

  echo $sql;

  //need to deal with cases here - we can do this later:
  // 1. The record is created successfully
  // 2. The record already exists in the database
  // 3. The record was not created in the database

  if ($conn->query($sql) === TRUE) {
      echo "<p>New record created successfully</p>";

  } else {
      echo "<p>oh dear</p>";
  }

  $conn->close();

  //need to add response code here saying that the email has been added
  echo "<p>$email has been added</p>";
?>
