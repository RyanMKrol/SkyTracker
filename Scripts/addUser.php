<?php

  include 'credentials.php';

  $email = $_POST['addr'];

  // Create connection
  $conn = new mysqli($servername, $username, $password, $database);

  // Check connection
  if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
  }

  $email = $conn->real_escape_string($email);

  $sql = "INSERT INTO Users (UserEmailAddress) VALUES ('$email');";

  //need to deal with cases here - we can do this later:
  // 1. The record is created successfully, in which case we just welcome the user to the platform
  // 2. The record already exists in the database, in which case we update the user's details, and give confirmation of that
  // 3. The record was not created in the database, we report an error

  if ($conn->query($sql) === TRUE) {
      echo "<p>New record created successfully</p>";

  } else {
      echo "<p>oh dear</p>";
  }

  $conn->close();

  //need to add response code here saying that the email has been added
  echo "<p>$email has been added</p>";
?>
