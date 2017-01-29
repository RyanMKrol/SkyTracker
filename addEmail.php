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

  $sql = "INSERT INTO Emails (EmailAddress) VALUES (\"$email\");";

  if ($conn->query($sql) === TRUE) {
      echo "<p>New record created successfully</p>";
  } else {
      echo "<p>oh dear</p>";
  }

  $conn->close();

  //need to add response code here saying that the email has been added
  echo "<p>$email has been added</p>";
?>
